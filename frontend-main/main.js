function showSection(sectionId) {
  document.querySelectorAll(".section").forEach((section) => {
    section.classList.remove("active");
  });

  if (sectionId === 'profile') {
    loadProfile(); 
  }

  document.getElementById(sectionId).classList.add("active");

  document.querySelectorAll("nav button").forEach((btn) => {
    btn.classList.remove("active");
  });
  const activeBtn = document.querySelector(`#btn-${sectionId}`);
  if (activeBtn) activeBtn.classList.add("active");

  localStorage.setItem("activeSection", sectionId);
}

async function loadProfile() {
  const token = localStorage.getItem('token');
  if (!token) return;

  try {
    const res = await fetch('/api/profile', {
      method: 'GET',
      headers: { 'Authorization': 'Bearer ' + token }
    });

    if (res.ok) {
      const data = await res.json();
      document.getElementById('full_name').value = data.full_name || '';
      document.getElementById('email').value = data.email || '';
      document.getElementById('street').value = data.street || '';
      document.getElementById('house').value = data.house || '';
      document.getElementById('apartment').value = data.apartment || '';
    } else {
      document.getElementById('profile-message').textContent = 'Failed to load profile';
    }
  } catch (err) {
    console.error(err);
    document.getElementById('profile-message').textContent = 'Error loading profile';
  }
}

function initLogout() {
  const logoutBtn = document.getElementById("logout-btn");
  if (logoutBtn) {
    logoutBtn.addEventListener("click", () => {
      localStorage.removeItem("token");
      window.location.href = "login.html";
    });
  }
}

document.addEventListener("DOMContentLoaded", () => {
  console.log("=== PAGE LOADED ===");
  
  const token = localStorage.getItem("token");
  console.log("Token exists:", !!token);
  console.log("Token value:", token ? token.substring(0, 20) + "..." : "null");
  
  if (!token) {
    console.log("No token found, redirecting to login");
    window.location.href = "login.html";
    return;
  }
  
  console.log("Token found, staying on main page");

  document.getElementById("username").textContent = "Test User";
  //document.getElementById("profile-info").innerHTML = "<p>Profile loaded successfully (test mode)</p>";
  document.getElementById("report-list").innerHTML = "<p>Reports section (test mode)</p>";
  document.getElementById("news-list").innerHTML = "<p>News section (test mode)</p>";
  document.getElementById("notifications-list").innerHTML = "<p>Notifications section (test mode)</p>";

  showSection("profile");
  
  initLogout();
});

const profileForm = document.getElementById('profile-form');
if (profileForm) {
  profileForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const token = localStorage.getItem('token');
    const msg = document.getElementById('profile-message');
    msg.style.display = 'block';
    msg.textContent = 'Updating...';

    const body = {
      full_name: document.getElementById('full_name').value,
      street: document.getElementById('street').value,
      house: document.getElementById('house').value,
      apartment: document.getElementById('apartment').value
    };

    try {
      const res = await fetch('/api/update-profile', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ' + token
        },
        body: JSON.stringify(body)
      });

      if (res.ok) {
        msg.textContent = 'Profile updated successfully!';
        msg.className = 'message success';
      } else {
        const error = await res.text();
        msg.textContent = 'Update failed: ' + error;
        msg.className = 'message error';
      }
    } catch (err) {
      msg.textContent = 'Error: ' + err.message;
      msg.className = 'message error';
    }
  });
}
