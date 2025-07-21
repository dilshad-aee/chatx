function signup() {
    const username = document.getElementById("username").value.trim();
    const password = document.getElementById("password").value.trim();
    const repassword = document.getElementById("repassword").value.trim();
    const log = document.getElementById("log");

    if (!username || !password || !repassword) {
        log.style.color = "red";
        log.textContent = " Fields cannot be blank";
        return;
    }
    if(password.length < 8) {
        log.style.color = "red";
        log.textContent = "Password must be at least 8 characters long";
        return;
    }
    

    if (repassword !== password) {
        log.style.color = "red";
        log.textContent = "Passwords do not match";
        return;
    }

    fetch("http://localhost:8080/signup", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ username, password })
    })
    .then(server_response => {
        if (!server_response.ok) {
            return server_response.text().then(text => { throw new Error(text); });
        }
        return server_response.json();
    })
    .then(data => {
        log.style.color = "green";
        log.textContent = "✅ Signed up as: " + data.username;
    })
    .catch(err => {
        log.style.color = "red";
        log.textContent = "❌ " + err.message;
    });
}