// DOM Elements
const documentsSection = document.getElementById('documents');
const documentsList = document.getElementById('documents-list');
const documentFilter = document.getElementById('document-filter');
const uploadBtn = document.getElementById('btn-upload');
const startUploadBtn = document.getElementById('start-upload');
const uploadModal = document.getElementById('upload-modal');
const closeModal = document.querySelector('.close');
const fileInput = document.getElementById('document-file');
const fileInfo = document.getElementById('file-info');
const dropZone = document.getElementById('drop-zone');
const uploadProgress = document.getElementById('upload-progress');
const progressBar = document.querySelector('.progress');
const progressText = document.querySelector('.progress-text');
const uploadStatus = document.getElementById('upload-status');

// Configuration
const API_BASE_URL = 'http://localhost:50051';
const USER_ID = 'demo-user'; // Replace with actual user ID in real system

// Initialize the application
function init() {
    loadDocuments();
    setupEventListeners();
}

// Load documents from the server
function loadDocuments() {
    showLoader(documentsList);
    
    // In a real system: fetch(`${API_BASE_URL}?userId=${USER_ID}`)
    // For demo, we'll simulate API call
    setTimeout(() => {
        const demoDocuments = [
            {
                id: "doc1",
                filename: "Project Proposal.pdf",
                type: "application/pdf",
                createdAt: "2025-05-01T10:30:00Z",
                size: 2457621
            },
            {
                id: "doc2",
                filename: "Financial Report.xlsx",
                type: "application/vnd.ms-excel",
                createdAt: "2025-05-10T14:45:00Z",
                size: 1876543
            },
            {
                id: "doc3",
                filename: "Contract Agreement.docx",
                type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
                createdAt: "2025-04-15T09:15:00Z",
                size: 3210987
            },
            {
                id: "doc4",
                filename: "Technical Specifications.docx",
                type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
                createdAt: "2025-05-18T11:20:00Z",
                size: 1567890
            },
            {
                id: "doc5",
                filename: "Meeting Notes.txt",
                type: "text/plain",
                createdAt: "2025-05-20T16:30:00Z",
                size: 87654
            },
            {
                id: "doc6",
                filename: "Design Mockups.png",
                type: "image/png",
                createdAt: "2025-05-22T09:45:00Z",
                size: 3456789
            }
        ];
        
        renderDocuments(demoDocuments);
    }, 1500);
}

// Render documents to the UI
function renderDocuments(documents) {
    documentsList.innerHTML = '';
    
    if (documents.length === 0) {
        documentsList.innerHTML = `
            <div class="empty-state">
                <i class="fas fa-folder-open"></i>
                <h3>No documents found</h3>
                <p>Upload your first document to get started</p>
                <button class="upload-btn" onclick="openUploadModal()">
                    <i class="fas fa-upload"></i> Upload Document
                </button>
            </div>
        `;
        return;
    }
    
    documents.forEach(doc => {
        const docElement = document.createElement('div');
        docElement.className = 'document-card';
        docElement.innerHTML = `
            <div class="document-icon">
                ${getFileIcon(doc.type)}
            </div>
            <div class="document-name">${doc.filename}</div>
            <div class="document-meta">
                <div>${formatFileType(doc.type)}</div>
                <div>${formatDate(doc.createdAt)} • ${formatFileSize(doc.size)}</div>
            </div>
            <div class="document-actions">
                <button class="btn-download" onclick="downloadDocument('${doc.id}')">
                    <i class="fas fa-download"></i> Download
                </button>
                <button class="btn-delete" onclick="deleteDocument('${doc.id}')">
                    <i class="fas fa-trash"></i> Delete
                </button>
            </div>
        `;
        documentsList.appendChild(docElement);
    });
}

// Open upload modal
function openUploadModal() {
    uploadModal.style.display = 'flex';
    resetUploadForm();
}

// Close upload modal
function closeUploadModal() {
    uploadModal.style.display = 'none';
    resetUploadForm();
}

// Reset upload form
function resetUploadForm() {
    fileInput.value = '';
    fileInfo.innerHTML = '';
    uploadProgress.style.display = 'none';
    uploadStatus.innerHTML = '';
    uploadStatus.className = '';
    progressBar.style.width = '0%';
    progressText.textContent = '0%';
}

// Setup event listeners
function setupEventListeners() {
    // Upload button
    uploadBtn.addEventListener('click', openUploadModal);
    
    // Close modal
    closeModal.addEventListener('click', closeUploadModal);
    
    // File input change
    fileInput.addEventListener('change', handleFileSelect);
    
    // Start upload
    startUploadBtn.addEventListener('click', uploadDocument);
    
    // Document search
    documentFilter.addEventListener('input', filterDocuments);
    
    // Drag and drop events
    dropZone.addEventListener('dragover', handleDragOver);
    dropZone.addEventListener('dragleave', handleDragLeave);
    dropZone.addEventListener('drop', handleDrop);
    
    // Close modal when clicking outside
    window.addEventListener('click', (e) => {
        if (e.target === uploadModal) {
            closeUploadModal();
        }
    });
}

// Handle file selection
function handleFileSelect(e) {
    const file = e.target.files[0];
    if (file) {
        displayFileInfo(file);
    }
}

// Display selected file info
function displayFileInfo(file) {
    fileInfo.innerHTML = `
        <div><strong>File:</strong> ${file.name}</div>
        <div><strong>Type:</strong> ${file.type || 'Unknown'}</div>
        <div><strong>Size:</strong> ${formatFileSize(file.size)}</div>
    `;
}

// Handle drag over
function handleDragOver(e) {
    e.preventDefault();
    dropZone.classList.add('dragover');
}

// Handle drag leave
function handleDragLeave(e) {
    e.preventDefault();
    dropZone.classList.remove('dragover');
}

// Handle file drop
function handleDrop(e) {
    e.preventDefault();
    dropZone.classList.remove('dragover');
    
    const file = e.dataTransfer.files[0];
    if (file) {
        fileInput.files = e.dataTransfer.files;
        displayFileInfo(file);
    }
}

// Upload document
function uploadDocument() {
    const file = fileInput.files[0];
    if (!file) {
        showToast('Please select a file first', 'error');
        return;
    }
    
    // Show upload progress
    uploadProgress.style.display = 'block';
    startUploadBtn.disabled = true;
    
    // Simulate upload progress
    let progress = 0;
    const interval = setInterval(() => {
        progress += Math.floor(Math.random() * 10) + 1;
        if (progress >= 100) {
            progress = 100;
            clearInterval(interval);
            
            // Upload complete
            setTimeout(() => {
                uploadStatus.innerHTML = `
                    <div class="status-icon"><i class="fas fa-check-circle"></i></div>
                    <div class="status-message">
                        <h4>Upload Successful!</h4>
                        <p>Your document has been uploaded successfully.</p>
                    </div>
                `;
                uploadStatus.className = 'upload-success';
                
                // Add new document to the list
                const newDocument = {
                    id: `doc${Date.now()}`,
                    filename: file.name,
                    type: file.type || 'application/octet-stream',
                    createdAt: new Date().toISOString(),
                    size: file.size
                };
                
                // Reload documents after a delay
                setTimeout(() => {
                    closeUploadModal();
                    loadDocuments();
                    showToast('Document uploaded successfully', 'success');
                }, 2000);
            }, 500);
        }
        
        // Update progress bar
        progressBar.style.width = `${progress}%`;
        progressText.textContent = `${progress}%`;
    }, 200);
}

// Download document
function downloadDocument(documentId) {
    showToast('Downloading document...');
    
    // In a real system: fetch(`${API_BASE_URL}/${documentId}/download`)
    setTimeout(() => {
        // Create a dummy file to download
        const filename = `document-${documentId}.txt`;
        const content = `This is a demo file for document ID: ${documentId}\n\n` +
                        `This file was downloaded from the Document Management System.\n` +
                        `In a real system, this would be your actual document content.`;
        
        const blob = new Blob([content], { type: 'text/plain' });
        const url = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = filename;
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
        URL.revokeObjectURL(url);
        
        showToast('Document downloaded', 'success');
    }, 1500);
}

// Delete document
function deleteDocument(documentId) {
    if (!confirm('Are you sure you want to delete this document?')) return;
    
    showToast('Deleting document...');
    
    // In a real system: fetch(`${API_BASE_URL}/${documentId}`, { method: 'DELETE' })
    setTimeout(() => {
        // Remove document from UI
        loadDocuments();
        showToast('Document deleted', 'success');
    }, 1000);
}

// Filter documents
function filterDocuments() {
    const query = documentFilter.value.toLowerCase();
    const cards = documentsList.querySelectorAll('.document-card');
    
    cards.forEach(card => {
        const name = card.querySelector('.document-name').textContent.toLowerCase();
        card.style.display = name.includes(query) ? 'block' : 'none';
    });
}

// Helper functions
function formatDate(dateString) {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', { 
        year: 'numeric', 
        month: 'short', 
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    });
}

function formatFileSize(bytes) {
    if (bytes < 1024) return bytes + ' bytes';
    else if (bytes < 1048576) return (bytes / 1024).toFixed(1) + ' KB';
    else return (bytes / 1048576).toFixed(1) + ' MB';
}

function formatFileType(type) {
    const typeMap = {
        'application/pdf': 'PDF Document',
        'application/vnd.ms-excel': 'Excel Spreadsheet',
        'application/vnd.openxmlformats-officedocument.wordprocessingml.document': 'Word Document',
        'text/plain': 'Text File',
        'image/png': 'PNG Image',
        'image/jpeg': 'JPEG Image'
    };
    
    return typeMap[type] || type.split('/')[1].toUpperCase() + ' File';
}

function getFileIcon(type) {
    const iconMap = {
        'application/pdf': '<i class="fas fa-file-pdf"></i>',
        'application/vnd.ms-excel': '<i class="fas fa-file-excel"></i>',
        'application/vnd.openxmlformats-officedocument.wordprocessingml.document': '<i class="fas fa-file-word"></i>',
        'text/plain': '<i class="fas fa-file-alt"></i>',
        'image/png': '<i class="fas fa-file-image"></i>',
        'image/jpeg': '<i class="fas fa-file-image"></i>'
    };
    
    return iconMap[type] || '<i class="fas fa-file"></i>';
}

function showLoader(container) {
    container.innerHTML = '<div class="loader"></div>';
}

function showToast(message, type = 'info') {
    const toast = document.createElement('div');
    toast.className = `toast ${type}`;
    toast.innerHTML = `
        <i class="fas ${type === 'success' ? 'fa-check-circle' : type === 'error' ? 'fa-exclamation-circle' : 'fa-info-circle'}"></i>
        <span>${message}</span>
    `;
    
    const container = document.getElementById('toast-container');
    container.appendChild(toast);
    
    setTimeout(() => {
        toast.remove();
    }, 3000);
}

// Initialize the application when DOM is loadedы
document.addEventListener('DOMContentLoaded', init);