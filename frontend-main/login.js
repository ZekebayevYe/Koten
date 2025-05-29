document.addEventListener('DOMContentLoaded', () => {
  // Login
  const loginForm = document.getElementById('login-form');
  if (loginForm) {
    loginForm.addEventListener('submit', async (e) => {
      e.preventDefault();

      const email = loginForm.querySelector('input[type="email"]').value;
      const password = loginForm.querySelector('input[type="password"]').value;

      const res = await fetch('http://localhost:8080/api/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
      });

      if (res.ok) {
        const data = await res.json();
        localStorage.setItem('token', data.token);
        alert('Login successful!');
        window.location.href = 'index.html';
      } else {
        const error = await res.text();
        alert('Login failed: ' + error);
      }
    });
  }

  // Register
  const registerForm = document.getElementById('register-form');
  if (registerForm) {
    registerForm.addEventListener('submit', async (e) => {
      e.preventDefault();

      const name = registerForm.querySelector('input[type="text"]').value;
      const email = registerForm.querySelector('input[type="email"]').value;
      const password = registerForm.querySelector('input[type="password"]').value;

      const res = await fetch('http://localhost:8080/api/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, email, password })
      });

      if (res.ok) {
        alert('Registration successful! Please login.');
        // переключиться на логин форму
        document.getElementById('login-tab').click();
      } else {
        const error = await res.text();
        alert('Registration failed: ' + error);
      }
    });
  }
});
