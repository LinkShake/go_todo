const main = async () => {
  const userId = "stg";
  try {
    //select all necessary elements
    const todoContainer = document.getElementById("todo-container");
    const addTodoBtn = document.getElementById("btn-add-todo");
    const input = document.getElementById("todo-input");
    let deleteBtns = document.querySelectorAll(".btn-delete");
    //inputContent
    let inputContent = "";
    //fetch initial data
    const res = await fetch(`/todos/${userId}`);
    const data = await res.json();
    //display all todos associated to user
    if (Array.isArray(data)) {
      data.forEach((currTodo) => {
        const el = `
                <div id="todo-${currTodo.ID}">
                    <p>${currTodo.Text}</p>
                    <button id="delete-todo-${currTodo.ID}" class="btn-delete">delete</button>
                    <button id="edit-todo-${currTodo.ID}" class="btn-edit">edit</button>
                </div>
            `;
        todoContainer.insertAdjacentHTML("beforeend", el);
        deleteBtns = document.querySelectorAll(".btn-delete");
      });
    }
    //event handlers
    input.addEventListener("input", (e) => {
      inputContent += e.data;
    });
    addTodoBtn.addEventListener("click", async () => {
      try {
        if (!inputContent.length) return;
        const res = await fetch("/add-todo", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            userId,
            text: inputContent,
          }),
        });
        const data = await res.json();
        input.value = "";
        inputContent = "";
        const el = `
                <div id="todo-${data.ID}">
                    <p>${data.Text}</p>
                    <button id="delete-todo-${data.ID}" class="btn-delete">delete</button>
                    <button id="edit-todo-${data.ID}" class="btn-edit">edit</button>
                </div>
            `;
        todoContainer.insertAdjacentHTML("beforeend", el);
        deleteBtns = document.querySelectorAll(".btn-delete");
      } catch (err) {
        console.log(err);
      }
    });
    deleteBtns.forEach((currBtn) => {
      currBtn.addEventListener("click", async (e) => {
        try {
          const id = e.currentTarget.id.slice(
            e.currentTarget.id.lastIndexOf("-") + 1
          );
          const res = await fetch("/delete-todo", {
            method: "DELETE",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              userId,
              id: +id,
            }),
          });
          const data = await res.json();
          if (isNaN(data)) return;
          const deletedTodo = document.getElementById(`todo-${data}`);
          todoContainer.removeChild(deletedTodo);
        } catch (err) {
          console.log(err);
        }
      });
    });
  } catch (err) {
    console.log(err);
  }
};

main();
