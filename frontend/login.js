function login() {
  const username = document.getElementById("username").value.trim();
  const password = document.getElementById("password").value.trim();
  const log = document.getElementById("log");


  if (!username){
    log.textContent="Username cannot be blank"
    return;
  }

  if (!password){
    log.textContent="Password cannot be blank"
    return;
  }

  fetch("http://localhost:8080/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      username: username,
      password: password

    })
  })


  .then(res => {
    if (!res.ok) {
      return res.text().then(text => { throw new Error(text); });
    }
    return res.json();
  })
  .then(data => {

    localStorage.setItem("jwtToken",data.token)
    log.style.color = "green";
    log.textContent = "✅ Logged in as: " + data.username;

   
  })
  .catch(err => {
    log.style.color = "red";
    log.textContent = "❌ " + err.message;
  });
}