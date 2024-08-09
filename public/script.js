const createMarkup = (data) => {
  return `
    <div id="todo-${data.ID}">
      <p>${data.Text}</p>
      <input type="text" id="todo-${data.ID}-input" class="edit-todo-input hidden" min="1" max="20" />
      <button id="delete-todo-${data.ID}" class="btn-delete">delete</button>
      <button id="edit-todo-${data.ID}" class="btn-edit">edit</button>
      <button id="done-editing-todo-${data.ID}" class="btn-done-editing hidden">done</button>
    </div>
  `;
};

const toggleEditVisibility = (id) => {
  document
    .getElementById(`todo-${id}`)
    .querySelector("p")
    .classList.toggle("hidden");
  document.getElementById(`edit-todo-${id}`).classList.toggle("hidden");
  document.getElementById(`delete-todo-${id}`).classList.toggle("hidden");
  document.getElementById(`todo-${id}-input`).classList.toggle("hidden");
  document.getElementById(`done-editing-todo-${id}`).classList.toggle("hidden");
};

const main = async () => {
  const userId = "";
  try {
    //select all necessary elements
    const todoContainer = document.getElementById("todo-container");
    const addTodoBtn = document.getElementById("btn-add-todo");
    const input = document.getElementById("todo-input");
    let deleteBtns = document.querySelectorAll(".btn-delete");
    let editBtns = document.querySelectorAll(".btn-edit");
    let doneEditingBtns = document.querySelectorAll(".btn-done-editing");
    let editInputs = document.querySelectorAll(".edit-todo-input");
    //inputContent
    let inputContent = "";
    let editInputContent = "";
    //fetch initial data
    const res = await fetch(`/todos/${userId}`);
    const data = await res.json();
    //display all todos associated to user
    if (Array.isArray(data)) {
      data.forEach((currTodo) => {
        todoContainer.insertAdjacentHTML("beforeend", createMarkup(currTodo));
        deleteBtns = document.querySelectorAll(".btn-delete");
        editBtns = document.querySelectorAll(".btn-edit");
        doneEditingBtns = document.querySelectorAll(".btn-done-editing");
        editInputs = document.querySelectorAll(".edit-todo-input");
      });
    }
    //event handlers
    input.addEventListener("input", (e) => {
      if (e.data) inputContent += e.data;
    });
    input.addEventListener("change", (e) => {
      if (e.target.value) inputContent = e.target.value;
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
        todoContainer.insertAdjacentHTML("beforeend", createMarkup(data));
        deleteBtns = document.querySelectorAll(".btn-delete");
        editBtns = document.querySelectorAll(".btn-edit");
        doneEditingBtns = document.querySelectorAll(".btn-done-editing");
        editInputs = document.querySelectorAll(".edit-todo-input");
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
    editBtns.forEach((currBtn) => {
      currBtn.addEventListener("click", (e) => {
        try {
          const id = e.currentTarget.id.slice(
            e.currentTarget.id.lastIndexOf("-") + 1
          );
          editInputContent = document
            .getElementById(`todo-${id}`)
            .querySelector("p").textContent;
          toggleEditVisibility(id);
          document.getElementById(`todo-${id}-input`).value = editInputContent;
        } catch (err) {
          console.log(err);
        }
      });
    });
    doneEditingBtns.forEach((currBtn) => {
      currBtn.addEventListener("click", async (e) => {
        try {
          if (!editInputContent) return;
          const id = e.currentTarget.id.slice(
            e.currentTarget.id.lastIndexOf("-") + 1
          );
          const res = await fetch("/edit-todo", {
            method: "PUT",
            headers: {
              "Content-type": "application/json",
            },
            body: JSON.stringify({
              userId,
              id: +id,
              text: editInputContent,
            }),
          });
          const data = await res.json();
          editInputContent = "";
          document.getElementById(`todo-${id}`).querySelector("p").textContent =
            data.Text;
          toggleEditVisibility(id);
        } catch (err) {
          console.log(err);
        }
      });
    });
    editInputs.forEach((currInput) => {
      currInput.addEventListener("input", (e) => {
        if (e.data) editInputContent += e.data;
      });
      currInput.addEventListener("change", (e) => {
        if (e.target.value) editInputContent = e.target.value;
      });
    });
  } catch (err) {
    console.log(err);
  }
};

main();
