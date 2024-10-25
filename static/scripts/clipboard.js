// clipboard.js
export function copyToClipboard(filename) {
    fetch(`/icons/${filename}`)
        .then(response => response.text())
        .then(svgContent => {
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
