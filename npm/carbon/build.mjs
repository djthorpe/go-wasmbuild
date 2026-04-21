import { cpSync, existsSync, mkdirSync, readFileSync, rmSync, writeFileSync } from 'fs';
import { dirname, resolve } from 'path';
import { build } from 'esbuild';
import { fileURLToPath } from 'url';

const rootDir = dirname(fileURLToPath(import.meta.url));
const distDir = resolve(process.env.CARBON_DIST_DIR || rootDir);
const sourceAssetsDir = resolve(rootDir, 'assets');
const distAssetsDir = resolve(distDir, 'assets');
const sourceStylesPath = resolve(rootDir, 'node_modules', '@carbon', 'styles', 'css', 'styles.css');
const sourceIconsDir = resolve(rootDir, 'node_modules', '@carbon', 'icons', 'svg');

await import('./gen-icons.mjs');

mkdirSync(distDir, { recursive: true });

await build({
    entryPoints: [resolve(rootDir, 'index.js')],
    bundle: true,
    format: 'esm',
    minify: true,
    outfile: resolve(distDir, 'bundle.js'),
});

if (distDir !== rootDir) {
    rmSync(distAssetsDir, { recursive: true, force: true });
    cpSync(sourceAssetsDir, distAssetsDir, { recursive: true });
}

mkdirSync(distAssetsDir, { recursive: true });
rmSync(resolve(distAssetsDir, 'icons'), { recursive: true, force: true });
cpSync(sourceIconsDir, resolve(distAssetsDir, 'icons'), { recursive: true });

const stylesLines = readFileSync(sourceStylesPath, 'utf8').split('\n');
writeFileSync(
    resolve(distAssetsDir, 'themes.css'),
    '/* Carbon theme tokens (generated) */\n' + stylesLines.slice(2786, 4107).join('\n')
);
writeFileSync(
    resolve(distAssetsDir, 'grid.css'),
    '/* Carbon CSS grid (generated) */\n' + stylesLines.slice(1063, 1862).join('\n')
);

if (!existsSync(resolve(distDir, 'bundle.js'))) {
    throw new Error('bundle.js was not created');
}