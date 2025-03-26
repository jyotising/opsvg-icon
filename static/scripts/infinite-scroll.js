class IconLoader {
    constructor() {
        this.page = 1;
        this.loading = false;
        this.initialLoad = true;
        this.perPage = this.initialLoad ? 200 : 100;
        this.gallery = document.getElementById('iconGallery');
        this.spinner = document.getElementById('loadingSpinner');
        this.lastScrollPosition = 0;
        this.debounceTimer = null;
        
        this.init();
    }

    init() {
        // Load initial icons
        this.loadIcons();
        
        // Add scroll event listener
        window.addEventListener('scroll', () => {
            this.handleScroll();
        });
    }

    handleScroll() {
        // Clear existing timeout
        if (this.debounceTimer) clearTimeout(this.debounceTimer);
        
        // Debounce scroll event
        this.debounceTimer = setTimeout(() => {
            if (this.loading) return;

            const scrollPosition = window.innerHeight + window.scrollY;
            const contentHeight = document.documentElement.offsetHeight;
            
            // Load more when user is 200px from bottom
            if (scrollPosition >= contentHeight - 200) {
                this.perPage = 100; // Set to 100 for subsequent loads
                this.loadIcons();
            }
        }, 100);
    }

    async loadIcons() {
        try {
            this.loading = true;
            this.spinner.style.display = 'block';

            const response = await fetch(`/api/icons?page=${this.page}&per_page=${this.perPage}`);
            const data = await response.json();

            if (data.icons && data.icons.length > 0) {
                this.renderIcons(data.icons);
                this.page++;
            }

            this.initialLoad = false;
            this.loading = false;
            this.spinner.style.display = 'none';
        } catch (error) {
            console.error('Error loading icons:', error);
            this.loading = false;
            this.spinner.style.display = 'none';
        }
    }

    renderIcons(icons) {
        const iconHTML = icons.map(icon => `
            <div class="icon-item" onclick="showIconModal('${icon.filename}', '${icon.name}')">
                <div class="icon-preview">
                    <img src="/icons/${icon.filename}" alt="${icon.name}" class="svg-icon">
                </div>
                <p class="icon-name">${icon.name}</p>
            </div>
        `).join('');

        this.gallery.insertAdjacentHTML('beforeend', iconHTML);
    }
}

// Initialize the loader when DOM is ready
document.addEventListener('DOMContentLoaded', () => {
    new IconLoader();
}); 