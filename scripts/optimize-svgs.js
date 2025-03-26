const fs = require('fs');
const path = require('path');
const SVGO = require('svgo');

const svgo = new SVGO({
    plugins: [
        { removeViewBox: false }, // keep viewBox
        { removeDimensions: true }, // remove width/height
        { cleanupIDs: true },
        { removeUselessStrokeAndFill: true },
        { collapseGroups: true },
        { mergePaths: true },
        { convertPathData: true },
        { removeEmptyAttrs: true },
        { removeEmptyContainers: true },
        { removeUnusedNS: true },
        { removeTitle: true },
        { removeDesc: true },
        { removeXMLNS: false } // keep xmlns for standalone SVGs
    ]
});

const iconsDir = path.join(__dirname, '../icons');

async function optimizeSVGs() {
    const files = fs.readdirSync(iconsDir);
    
    for (const file of files) {
        if (path.extname(file) === '.svg') {
            const filePath = path.join(iconsDir, file);
            const svg = fs.readFileSync(filePath, 'utf8');
            
            try {
                const result = await svgo.optimize(svg, { path: filePath });
                fs.writeFileSync(filePath, result.data);
                console.log(`✓ Optimized: ${file}`);
            } catch (error) {
                console.error(`✗ Error optimizing ${file}:`, error);
            }
        }
    }
}

optimizeSVGs().then(() => {
    console.log('SVG optimization complete!');
}); 