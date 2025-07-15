async function register() {
  const fullname = document.getElementById("nameRegister").value;
  const username = document.getElementById("usernameRegister").value;
  const password = document.getElementById("passwordRegister").value;

  if (fullname === "") {
    alert("You have not entered your name");
    document.getElementById("nameRegister").focus();
    return;
  }

  if (username === "") {
    alert("You have not entered a username");
    document.getElementById("usernameRegister").focus();
    return;
  }

  if (password === "") {
    alert("You have not entered a password");
    document.getElementById("passwordRegister").focus();
    return;
  }

  console.log(fullname);

  try {
    const response = await fetch("http://localhost:8080/api/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        username,
        password,
        full_name: fullname,
      }),
    });

    if (response.ok) {
      alert("Registration successful\nPlease login again.");
      window.location.reload();
    } else {
      const result = await response.json();
      throw new Error(result.message || "Username already exists");
    }
  } catch (err) {
    alert(err.message);
  }
}
