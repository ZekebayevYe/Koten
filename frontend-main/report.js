const API_BASE = "http://localhost:8080/api/v1";

function showSection(sectionId) {
  document.querySelectorAll(".section").forEach((sec) => sec.classList.remove("active"));
  document.getElementById(sectionId).classList.add("active");

 
  if (sectionId === "report") loadReports();
  if (sectionId === "news") loadNews();
}

document.addEventListener("DOMContentLoaded", () => {
  loadReports();
  loadNews(); 

  document.getElementById("report-form").addEventListener("submit", async (e) => {
    e.preventDefault();
    const id = document.getElementById("report-id").value;
    const title = document.getElementById("report-title").value;
    const content = document.getElementById("report-content").value;

    const report = { title, content };

    if (id) {
      await fetch(`${API_BASE}/reports/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(report),
      });
    } else {
      await fetch(`${API_BASE}/reports`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(report),
      });
    }

    document.getElementById("report-form").reset();
    loadReports();
  });
});

async function loadReports() {
  const res = await fetch(`${API_BASE}/reports`);
  const reports = await res.json();
  const list = document.getElementById("report-list");
  list.innerHTML = "";

  reports.forEach((report) => {
    const div = document.createElement("div");
    div.className = "report-item";
    div.innerHTML = `
      <h3>${report.title}</h3>
      <p>${report.content}</p>
      <button onclick="editReport('${report.id}', '${report.title}', \`${report.content.replace(/`/g, "\\`")}\`)">Edit</button>
      <button onclick="deleteReport('${report.id}')">Delete</button>
    `;
    list.appendChild(div);
  });
}

function editReport(id, title, content) {
  document.getElementById("report-id").value = id;
  document.getElementById("report-title").value = title;
  document.getElementById("report-content").value = content;
  showSection("report");
}

async function deleteReport(id) {
  await fetch(`${API_BASE}/reports/${id}`, { method: "DELETE" });
  loadReports();
}

async function loadNews() {
  try {
    const res = await fetch(`${API_BASE}/news`);
    const news = await res.json();
    const list = document.getElementById("news-list");
    list.innerHTML = "";

    news.forEach((item) => {
      const div = document.createElement("div");
      div.className = "news-item";
      div.innerHTML = `
        <h3>${item.title}</h3>
        <p>${item.content}</p>
        <small>${new Date(item.created_at || item.date || Date.now()).toLocaleString()}</small>
        <hr />
      `;
      list.appendChild(div);
    });
  } catch (err) {
    document.getElementById("news-list").innerText = "Failed to load news.";
    console.error("Error loading news:", err);
  }
}
