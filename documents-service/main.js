// Демонстрационные данные для заполнения интерфейса
const demoData = {
  reports: [
    { title: "Monthly Water Usage", status: "Completed" },
    { title: "Electricity Consumption", status: "Pending" },
    { title: "Gas Usage Analysis", status: "In Progress" }
  ],
  news: [
    { date: "2025-05-15", text: "New water conservation initiative launched" },
    { date: "2025-05-10", text: "Scheduled maintenance next week" }
  ],
  profile: {
    name: "John Doe",
    email: "john.doe@example.com"
  }
};

// Основные функции
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
  
  // Загружаем документы при переходе на вкладку
  if (sectionId === 'documents') {
    fetchDocuments();
  }
}

function showLoader(containerId) {
  const container = document.getElementById(containerId);
  container.innerHTML = '<div class="loader"></div>';
}

function showToast(message, isError = false) {
  const toast = document.createElement("div");
  toast.className = "toast";
  toast.textContent = message;
  
  if (isError) {
    toast.style.backgroundColor = "#dc3545";
  }
  
  document.getElementById("toast-container").appendChild(toast);
  setTimeout(() => toast.remove(), 3000);
}

function openUploadModal() {
  document.getElementById('upload-modal').style.display = 'block';
}

function closeUploadModal() {
  document.getElementById('upload-modal').style.display = 'none';
  document.getElementById('upload-status').innerHTML = '';
  document.getElementById('document-file').value = '';
  document.getElementById('file-name').textContent = 'No file selected';
}

function updateFileName() {
  const fileInput = document.getElementById('document-file');
  const fileName = document.getElementById('file-name');
  
  if (fileInput.files.length > 0) {
    fileName.textContent = fileInput.files[0].name;
  } else {
    fileName.textContent = 'No file selected';
  }
}

// Функции для работы с документами
function fetchDocuments() {
  showLoader('documents-list');
  
  // В демо-режиме используем задержку для имитации запроса
  setTimeout(() => {
    // Здесь будет реальный запрос к API
    // fetch('http://localhost:8080/api/documents', ...)
    
    // Демо-данные
    const demoDocuments = [
      {
        id: "doc1",
        filename: "Water Bill - May 2025.pdf",
        type: "application/pdf",
        createdAt: "2025-05-01T10:30:00Z"
      },
      {
        id: "doc2",
        filename: "Electricity Report.xlsx",
        type: "application/vnd.ms-excel",
        createdAt: "2025-05-10T14:45:00Z"
      },
      {
        id: "doc3",
        filename: "Apartment Contract.docx",
        type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
        createdAt: "2025-04-15T09:15:00Z"
      }
    ];
    
    renderDocuments(demoDocuments);
    showToast("Documents loaded");
  }, 1000);
}

function renderDocuments(documents) {
  const container = document.getElementById('documents-list');
  container.innerHTML = '';
  
  if (documents.length === 0) {
    container.innerHTML = '<p>No documents found. Upload your first document!</p>';
    return;
  }
  
  documents.forEach(doc => {
    const docElement = document.createElement('div');
    docElement.className = 'document-item';
    
    docElement.innerHTML = `
      <div class="document-info">
        <strong>${doc.filename}</strong>
        <div>${doc.type.split('/')[1].toUpperCase()} - ${new Date(doc.createdAt).toLocaleDateString()}</div>
      </div>
      <div class="document-actions">
        <button onclick="downloadDocument('${doc.id}')">
          <i class="fas fa-download"></i> Download
        </button>
        <button class="danger" onclick="deleteDocument('${doc.id}')">
          <i class="fas fa-trash"></i> Delete
        </button>
      </div>
    `;
    
    container.appendChild(docElement);
  });
}

function uploadDocument() {
  const fileInput = document.getElementById('document-file');
  
  if (!fileInput.files || fileInput.files.length === 0) {
    showToast("Please select a file", true);
    return;
  }
  
  const file = fileInput.files[0];
  showLoader('upload-status');
  
  // В демо-режиме используем задержку для имитации загрузки
  setTimeout(() => {
    // Здесь будет реальный запрос к API
    // fetch('http://localhost:8080/api/documents/upload', ...)
    
    // Демо-результат
    document.getElementById('upload-status').innerHTML = `
      <p style="color: green;">Document uploaded successfully!</p>
      <p>File: ${file.name}</p>
      <p>Size: ${Math.round(file.size / 1024)} KB</p>
    `;
    
    // Обновляем список документов
    fetchDocuments();
    showToast("Document uploaded");
    
    // Автоматически закрываем модальное окно через 2 секунды
    setTimeout(closeUploadModal, 2000);
  }, 1500);
}

function downloadDocument(documentId) {
  showToast("Downloading document...");
  
  // В демо-режиме используем задержку для имитации скачивания
  setTimeout(() => {
    // Здесь будет реальный запрос к API
    // fetch(`http://localhost:8080/api/documents/${documentId}/download`, ...)
    
    // Демо-скачивание
    showToast("Document downloaded");
    
    // Создаем временный файл для скачивания
    const blob = new Blob(["This is a demo file content"], { type: "text/plain" });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `document-${documentId}.txt`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  }, 1000);
}

function deleteDocument(documentId) {
  if (!confirm("Are you sure you want to delete this document?")) return;
  
  showToast("Deleting document...");
  
  // В демо-режиме используем задержку для имитации удаления
  setTimeout(() => {
    // Здесь будет реальный запрос к API
    // fetch(`http://localhost:8080/api/documents/${documentId}`, ...)
    
    // Обновляем список документов
    fetchDocuments();
    showToast("Document deleted");
  }, 800);
}

// Инициализация
document.addEventListener("DOMContentLoaded", () => {
  // Устанавливаем активную вкладку
  showSection('documents');
  
  // Заполняем демо-данные
  document.getElementById('report-list').innerHTML = 
    demoData.reports.map(r => `<p>${r.title} — ${r.status}</p>`).join("");
  
  document.getElementById('news-list').innerHTML = 
    demoData.news.map(n => `<p><strong>${n.date}</strong>: ${n.text}</p>`).join("");
  
  document.getElementById('profile-info').innerHTML = 
    `<p>Name: ${demoData.profile.name}</p><p>Email: ${demoData.profile.email}</p>`;
  
  document.getElementById('username').textContent = demoData.profile.name;
  
  // Загружаем документы
  fetchDocuments();
  
  // Инициализация фильтра документов
  const filterInput = document.getElementById('document-filter');
  filterInput.addEventListener('input', () => {
    const query = filterInput.value.toLowerCase();
    const items = document.querySelectorAll('.document-item');
    
    items.forEach(item => {
      const filename = item.querySelector('strong').textContent.toLowerCase();
      item.style.display = filename.includes(query) ? 'flex' : 'none';
    });
  });
});