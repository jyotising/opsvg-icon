 function copyToClipboard(filename) {
    // Fetch the SVG file using Fetch API
    fetch(`/icons/${filename}`)
        .then(response => response.text())
        .then(svgContent => {
            // Copy the SVG content to the clipboard
            navigator.clipboard.writeText(svgContent).then(() => {
                alert(`${filename} copied to clipboard!`);
            }).catch(err => {
                alert('Failed to copy SVG.');
                console.error(err);
            });
        })
        .catch(err => {
            alert('Error fetching SVG file.');
            console.error(err);
        });
}


// Listen to color picker changes


document.addEventListener("DOMContentLoaded", function() {
    const strokeColorPicker = document.getElementById("strokeColorPicker");
    const icons = document.querySelectorAll(".svg-icon");

    // Listen for changes on the stroke color picker
    strokeColorPicker.addEventListener("input", function() {
        const selectedStrokeColor = strokeColorPicker.value;

        // Update the stroke color of all SVG icons
        icons.forEach(icon => {
            fetch(icon.src)
                .then(response => response.text())
                .then(svgContent => {
                    // Modify the SVG content to set the stroke color
                    const coloredSVG = svgContent.replace(/stroke="[^"]*"/g, `stroke="${selectedStrokeColor}"`);
                    const blob = new Blob([coloredSVG], { type: 'image/svg+xml' });
                    const newUrl = URL.createObjectURL(blob);

                    // Update the icon's src to the new stroked SVG
                    icon.src = newUrl;
                });
        });
    });
});




// Download Button to Include the Selected Color

function downloadIcon(filename) {
    const strokeColor = document.getElementById("strokeColorPicker").value;

    fetch(`/icons/${filename}`)
        .then(response => response.text())
        .then(svgContent => {
            // Modify the SVG content to set the stroke color
            const coloredSVG = svgContent.replace(/stroke="[^"]*"/g, `stroke="${strokeColor}"`);
            const blob = new Blob([coloredSVG], { type: 'image/svg+xml' });
            const url = URL.createObjectURL(blob);

            // Create a temporary link to trigger download
            const a = document.createElement("a");
            a.href = url;
            a.download = filename;
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
        })
        .catch(err => {
            console.error('Failed to download icon:', err);
        });
}

