// main.js

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

function showLoader(containerId) {
  const container = document.getElementById(containerId);
  container.innerHTML = '<div class="loader"></div>';
}

function showToast(message) {
  const toast = document.createElement("div");
  toast.className = "toast";
  toast.textContent = message;
  document.getElementById("toast-container").appendChild(toast);
  setTimeout(() => toast.remove(), 3000);
}

function fetchReports() {
  showLoader("report-list");
  fetch("http://localhost:8080/api/reports")
    .then((res) => res.json())
    .then((data) => {
      window._reportsData = data;
      renderReports(data);
      showToast("Reports loaded");
    })
    .catch(() => {
      document.getElementById("report-list").textContent = "Failed to load reports";
    });
}

function renderReports(data) {
  const container = document.getElementById("report-list");
  container.innerHTML = data.map(r => `<p>${r.title} — ${r.status}</p>`).join("");
}

function fetchNews() {
  showLoader("news-list");
  fetch("http://localhost:8080/api/news")
    .then((res) => res.json())
    .then((data) => {
      document.getElementById("news-list").innerHTML = data.map(n => `<p><strong>${n.date}</strong>: ${n.text}</p>`).join("");
      showToast("News loaded");
    })
    .catch(() => {
      document.getElementById("news-list").textContent = "Failed to load news";
    });
}

function fetchNotifications() {
  showLoader("notifications-list");
  fetch("http://localhost:8080/api/notifications")
    .then((res) => res.json())
    .then((data) => {
      document.getElementById("notifications-list").innerHTML = data.map(n => `<p>${n.message}</p>`).join("");
      showToast("Notifications loaded");
    })
    .catch(() => {
      document.getElementById("notifications-list").textContent = "Failed to load notifications";
    });
}

function fetchProfile() {
  showLoader("profile-info");
  fetch("http://localhost:8080/api/profile")
    .then((res) => res.json())
    .then((data) => {
      document.getElementById("profile-info").innerHTML = `<p>Name: ${data.name}</p><p>Email: ${data.email}</p>`;
      document.getElementById("username").textContent = data.name;
      showToast("Profile loaded");
    })
    .catch(() => {
      document.getElementById("profile-info").textContent = "Failed to load profile";
    });
}

function initThemeToggle() {
  const toggleBtn = document.createElement("button");
  toggleBtn.textContent = "🌙 Toggle Theme";
  toggleBtn.id = "themeToggleBtn";
  toggleBtn.style.cssText = "position:fixed;bottom:20px;right:20px;padding:10px;border:none;background:#333;color:#fff;border-radius:6px;cursor:pointer;z-index:1000";
  document.body.appendChild(toggleBtn);

  const isDark = localStorage.getItem("darkMode") === "true";
  if (isDark) document.body.classList.add("dark-mode");

  toggleBtn.addEventListener("click", () => {
    document.body.classList.toggle("dark-mode");
    localStorage.setItem("darkMode", document.body.classList.contains("dark-mode"));
  });
}

function initReportFilter() {
  const filterInput = document.getElementById("report-filter");
  filterInput.addEventListener("input", () => {
    const query = filterInput.value.toLowerCase();
    const filtered = window._reportsData.filter(r => r.title.toLowerCase().includes(query));
    renderReports(filtered);
  });
}

document.addEventListener("DOMContentLoaded", () => {
  const lastSection = localStorage.getItem("activeSection") || "report";
  showSection(lastSection);
  fetchReports();
  fetchNews();
  fetchNotifications();
  fetchProfile();
  initThemeToggle();
  initReportFilter();
});
