// Carbon Web Components — add new component imports here
import 'https://esm.sh/@carbon/web-components@2/es/components/ui-shell/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/code-snippet/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/button/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/tag/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/notification/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/accordion/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/tabs/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/text-input/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/password-input/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/search/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/number-input/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/textarea/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/select/index.js';
import 'https://esm.sh/@carbon/web-components@2/es/components/dropdown/index.js';

// ── cds-icon ──────────────────────────────────────────────────────────────────
// Lightweight Carbon icon renderer.  Loads SVG data on demand from the
// @carbon/icons package via esm.sh (one fetch per name+size, cached).
//
// Usage:  <cds-icon name="add"></cds-icon>
//         <cds-icon name="warning" size="24"></cds-icon>
//
// Color follows the CSS `color` property (fill:currentColor), so you can
// tint icons with any CSS color or Carbon token.
const _iconCache = new Map();

async function _loadIcon(name, size) {
    const key = `${name}/${size}`;
    if (_iconCache.has(key)) return _iconCache.get(key);
    try {
        const mod = await import(`https://esm.sh/@carbon/icons@11/es/${name}/${size}.js`);
        _iconCache.set(key, mod.default);
        return mod.default;
    } catch {
        return null;
    }
}

function _buildSVG(data) {
    const ns = 'http://www.w3.org/2000/svg';
    const svg = document.createElementNS(ns, 'svg');
    for (const [k, v] of Object.entries(data.attrs || {})) svg.setAttribute(k, String(v));
    svg.setAttribute('fill', 'currentColor');
    svg.setAttribute('aria-hidden', 'true');
    svg.style.display = 'inline-block';
    svg.style.verticalAlign = 'middle';
    for (const child of (data.content || [])) {
        const el = document.createElementNS(ns, child.elem);
        for (const [k, v] of Object.entries(child.attrs || {})) el.setAttribute(k, String(v));
        svg.appendChild(el);
    }
    return svg;
}

class CdsIconElement extends HTMLElement {
    static observedAttributes = ['name', 'size'];
    connectedCallback() { this._render(); }
    attributeChangedCallback() { if (this.isConnected) this._render(); }
    async _render() {
        const name = this.getAttribute('name');
        const size = this.getAttribute('size') || '16';
        if (!name) return;
        const data = await _loadIcon(name, size);
        if (data) { this.innerHTML = ''; this.appendChild(_buildSVG(data)); }
    }
}

customElements.define('cds-icon', CdsIconElement);
