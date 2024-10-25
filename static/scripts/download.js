// download.js
export function downloadIcon(filename) {
    const strokeColor = document.getElementById("strokeColorPicker").value;

    fetch(`/icons/${filename}`)
        .then(response => response.text())
        .then(svgContent => {
            const coloredSVG = svgContent.replace(/stroke="[^"]*"/g, `stroke="${strokeColor}"`);
            const blob = new Blob([coloredSVG], { type: 'image/svg+xml' });
            const url = URL.createObjectURL(blob);

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
