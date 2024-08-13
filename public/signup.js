const main = async () => {
  try {
    const res = await fetch("/is-user-logged-in");
    console.log(res);
    if (res.url === "http://localhost:3000/") {
      window.location.href = res.url;
      return;
    }
    const emailInput = document.getElementById("email");
    const pwdInput = document.getElementById("pwd");
    const btn = document.querySelector("button");
    let emailInputCont = "";
    let pwdInputCont = "";
    emailInput.addEventListener("input", (e) => {
      if (e.data) emailInputCont += e.data;
    });
    emailInput.addEventListener("change", (e) => {
      if (e.target.value) emailInputCont = e.target.value;
    });
    pwdInput.addEventListener("input", (e) => {
      if (e.data) pwdInputCont += e.data;
    });
    pwdInput.addEventListener("change", (e) => {
      if (e.target.value) pwdInputCont = e.target.value;
    });
    btn.addEventListener("click", async (e) => {
      e.preventDefault();
      if (pwdInputCont && emailInputCont) {
        const res = await fetch("/signup", {
          method: "POST",
          headers: {
            "Content-type": "application/json",
          },
          body: JSON.stringify({
            email: emailInputCont,
            pwd: pwdInputCont,
          }),
        });
        if (res.url === "http://localhost:3000/") {
          window.location.href = res.url;
        }
        const data = await res.json();
        if (!data.Ok) {
          alert(data.Msg);
        }
        // if (res.url) {
        //   window.location.replace(res.url);
        // }
      }
    });
  } catch (err) {
    console.log(err);
  }
};

main();
