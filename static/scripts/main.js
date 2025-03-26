// main.js

// Import the functions from other files
import { copyToClipboard } from './clipboard.js';
import { setupColorChange } from './colorChange.js';
import { downloadIcon } from './download.js';

// Initialize the event listeners or any setup needed
document.addEventListener("DOMContentLoaded", function() {
    // Setup color change listener
    setupColorChange();

    // You can also add additional setup code here if needed
});

// Make sure to expose the copyToClipboard and downloadIcon functions globally
window.copyToClipboard = copyToClipboard;
window.downloadIcon = downloadIcon;

function showIconModal(filename, name) {
    const modal = document.getElementById('iconModal');
    const modalTitle = modal.querySelector('.modal-title');
    const modalIcon = modal.querySelector('.modal-icon');
    const downloadBtn = document.getElementById('modalDownloadBtn');
    const copyBtn = document.getElementById('modalCopyBtn');
    
    modalTitle.textContent = name;
    modalIcon.src = `/icons/${filename}`;
    modalIcon.alt = name;
    
    downloadBtn.onclick = () => downloadIcon(filename);
    copyBtn.onclick = () => copyToClipboard(filename);
    
    modal.classList.add('active');
    
    // Close modal when clicking overlay or close button
    const closeModal = () => modal.classList.remove('active');
    modal.querySelector('.modal-overlay').onclick = closeModal;
    modal.querySelector('.modal-close').onclick = closeModal;
}

// Close modal on escape key
document.addEventListener('keydown', (e) => {
    if (e.key === 'Escape') {
        document.getElementById('iconModal').classList.remove('active');
    }
});
