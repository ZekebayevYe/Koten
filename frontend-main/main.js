function showSection(sectionId) {
  document.querySelectorAll(".section").forEach((section) => {
    section.classList.remove("active");
  });
  document.getElementById(sectionId).classList.add("active");

  document.querySelectorAll("nav button").forEach((btn) => {
    btn.classList.remove("active");
  });
  const activeBtn = document.querySelector(`#btn-${sectionId}`);
  if (activeBtn) activeBtn.classList.add("active");

  localStorage.setItem("activeSection", sectionId);
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
  document.getElementById("profile-info").innerHTML = "<p>Profile loaded successfully (test mode)</p>";
  document.getElementById("report-list").innerHTML = "<p>Reports section (test mode)</p>";
  document.getElementById("news-list").innerHTML = "<p>News section (test mode)</p>";
  document.getElementById("notifications-list").innerHTML = "<p>Notifications section (test mode)</p>";

  showSection("profile");
  
  initLogout();
});