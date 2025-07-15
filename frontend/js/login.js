async function login() {
  const username = document.getElementById("usernameLogin").value;
  const password = document.getElementById("passwordLogin").value;

  if (username === "") {
    alert("You have not entered a username");
    document.getElementById("usernameLogin").focus();
    return;
  }

  if (password === "") {
    alert("You have not entered a password");
    document.getElementById("passwordLogin").focus();
    return;
  }

  console.log(username);

  try {
    const response = await fetch("http://localhost:8080/api/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username, password }),
    });

    if (response.ok) {
      alert("Đăng nhập thành công");
      window.location.href = "dashboard.html";
    } else {
      const result = await response.json();
      throw new Error(result.message || "Incorrect username or password!");
    }
  } catch (err) {
    alert(err.message);
  }
}
