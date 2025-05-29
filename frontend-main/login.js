document.addEventListener('DOMContentLoaded', () => {
  // Login
  const loginForm = document.getElementById('login-form');
  if (loginForm) {
    loginForm.addEventListener('submit', async (e) => {
      e.preventDefault();
      
      const msgElement = document.getElementById("login-message");
      msgElement.textContent = "Logging in...";
      msgElement.className = "message";
      msgElement.style.display = "block";

      try {
        const email = loginForm.querySelector('input[type="email"]').value;
        const password = loginForm.querySelector('input[type="password"]').value;

        // 🔥 ИСПРАВЛЕНО: Используем относительный URL
        const res = await fetch('/api/login', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ email, password })
        });

        if (res.ok) {
          const data = await res.json();
          localStorage.setItem('token', data.token);
          msgElement.textContent = "Login successful!";
          msgElement.className = "message success";
          setTimeout(() => {
            // 🔥 ИСПРАВЛЕНО: Переход на главную без .html
            window.location.href = '/';
          }, 1000);
        } else {
          const error = await res.text();
          msgElement.textContent = 'Login failed: ' + error;
          msgElement.className = "message error";
        }
      } catch (err) {
        msgElement.textContent = 'Network error: ' + err.message;
        msgElement.className = "message error";
      }
    });
  }

  // Register
  const registerForm = document.getElementById('register-form');
  if (registerForm) {
    registerForm.addEventListener('submit', async (e) => {
      e.preventDefault();
      
      const msgElement = document.getElementById("register-message");
      msgElement.textContent = "Registering...";
      msgElement.className = "message";
      msgElement.style.display = "block";

      try {
        const name = registerForm.querySelector('input[type="text"]').value;
        const email = registerForm.querySelector('input[type="email"]').value;
        const password = registerForm.querySelector('input[type="password"]').value;

        // 🔥 ПРОБЛЕМА: Backend ожидает full_name, а не name
        const requestBody = {
          email: email,
          password: password,
          full_name: name, // 🔥 ИСПРАВЛЕНО: используем full_name
          house: "",       // 🔥 ДОБАВЛЕНО: обязательные поля
          street: "",      // 🔥 ДОБАВЛЕНО: обязательные поля  
          apartment: ""    // 🔥 ДОБАВЛЕНО: обязательные поля
        };

        // 🔥 ИСПРАВЛЕНО: Используем относительный URL
        const res = await fetch('/api/register', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(requestBody)
        });

        if (res.ok) {
          const data = await res.json();
          msgElement.textContent = 'Registration successful! Please login.';
          msgElement.className = "message success";
          setTimeout(() => {
            document.getElementById('login-tab').click();
            msgElement.style.display = "none";
            // Очищаем форму
            registerForm.reset();
          }, 2000);
        } else {
          const error = await res.text();
          msgElement.textContent = 'Registration failed: ' + error;
          msgElement.className = "message error";
        }
      } catch (err) {
        msgElement.textContent = 'Network error: ' + err.message;
        msgElement.className = "message error";
      }
    });
  }
});