document.addEventListener('DOMContentLoaded', function() {
    const searchInput = document.getElementById('searchInput');
    const searchForm = document.getElementById('searchForm');
    const iconGallery = document.querySelector('.icon-gallery');
    let debounceTimer;

    // Function to update URL without reloading the page
    function updateURL(searchTerm) {
        const url = new URL(window.location);
        if (searchTerm) {
            url.searchParams.set('search', searchTerm);
        } else {
            url.searchParams.delete('search');
        }
        window.history.pushState({}, '', url);
    }

    // Function to show loading state
    function setLoadingState(isLoading) {
        if (isLoading) {
            iconGallery.classList.add('loading');
            // Add loading indicator
            const loader = document.createElement('div');
            loader.className = 'search-loader';
            iconGallery.appendChild(loader);
        } else {
            iconGallery.classList.remove('loading');
            const loader = iconGallery.querySelector('.search-loader');
            if (loader) loader.remove();
        }
    }

    // Function to fetch and update icons
    async function fetchIcons(searchTerm) {
        try {
            setLoadingState(true);
            const response = await fetch(`/?search=${encodeURIComponent(searchTerm)}`);
            const text = await response.text();
            
            // Create a temporary element to parse the HTML
            const parser = new DOMParser();
            const doc = parser.parseFromString(text, 'text/html');
            const newGallery = doc.querySelector('.icon-gallery');
            
            // Update the gallery content
            iconGallery.innerHTML = newGallery.innerHTML;
            
            // Update URL
            updateURL(searchTerm);
        } catch (error) {
            console.error('Error fetching icons:', error);
        } finally {
            setLoadingState(false);
        }
    }

    // Handle real-time search with debouncing
    searchInput.addEventListener('input', (e) => {
        clearTimeout(debounceTimer);
        const searchTerm = e.target.value.trim();
        
        debounceTimer = setTimeout(() => {
            fetchIcons(searchTerm);
        }, 300); // Wait 300ms after user stops typing
    });

    // Prevent form submission (we're handling it via AJAX)
    searchForm.addEventListener('submit', (e) => {
        e.preventDefault();
    });

    // Handle keyboard shortcuts
    document.addEventListener('keydown', (e) => {
        // Command/Ctrl + K to focus search
        if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
            e.preventDefault();
            searchInput.focus();
        }
    });
});