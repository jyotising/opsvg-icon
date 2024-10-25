// colorChange.js
export function setupColorChange() {
    const strokeColorPicker = document.getElementById("strokeColorPicker");
    const icons = document.querySelectorAll(".svg-icon");

    strokeColorPicker.addEventListener("input", function() {
        const selectedStrokeColor = strokeColorPicker.value;

        icons.forEach(icon => {
            fetch(icon.src)
                .then(response => response.text())
                .then(svgContent => {
                    const coloredSVG = svgContent.replace(/stroke="[^"]*"/g, `stroke="${selectedStrokeColor}"`);
                    const blob = new Blob([coloredSVG], { type: 'image/svg+xml' });
                    const newUrl = URL.createObjectURL(blob);
                    icon.src = newUrl;
                });
        });
    });
}
