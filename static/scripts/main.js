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
