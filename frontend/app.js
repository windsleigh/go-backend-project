let authToken = "";

// Handle login
document.getElementById("loginForm").addEventListener("submit", async (e) => {
    e.preventDefault();
    const username = document.getElementById("username").value;

    try {
        const response = await fetch("http://localhost:8080/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ username }),
        });
        const data = await response.json();

        if (response.ok) {
            authToken = data.token;
            document.getElementById("loginMessage").textContent = "Login successful!";
        } else {
            document.getElementById("loginMessage").textContent = `Login failed: ${data.error}`;
        }
    } catch (error) {
        document.getElementById("loginMessage").textContent = `Error: ${error.message}`;
    }
});

// Handle create user
document.getElementById("createUserForm").addEventListener("submit", async (e) => {
    e.preventDefault();
    const name = document.getElementById("name").value;
    const email = document.getElementById("email").value;

    try {
        const response = await fetch("http://localhost:8080/users/create", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${authToken}`,
            },
            body: JSON.stringify({ name, email }),
        });
        const data = await response.json();

        if (response.ok) {
            document.getElementById("createUserMessage").textContent = data.message;
        } else {
            document.getElementById("createUserMessage").textContent = `Error: ${data.error}`;
        }
    } catch (error) {
        document.getElementById("createUserMessage").textContent = `Error: ${error.message}`;
    }
});

// Handle get users
document.getElementById("getUsersButton").addEventListener("click", async () => {
    try {
        const response = await fetch("http://localhost:8080/users", {
            headers: { Authorization: `Bearer ${authToken}` },
        });
        const users = await response.json();

        const userList = document.getElementById("userList");
        userList.innerHTML = ""; // Clear previous entries
        users.forEach((user) => {
            const li = document.createElement("li");
            li.textContent = `${user.name} (${user.email})`;
            userList.appendChild(li);
        });
    } catch (error) {
        document.getElementById("userList").textContent = `Error: ${error.message}`;
    }
});
