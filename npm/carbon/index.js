// Carbon Web Components — add or remove component imports here.
// After editing, run `make npm/carbon` from the repo root to rebuild bundle.js.

import '@carbon/web-components/es/components/ui-shell/index.js';
import '@carbon/web-components/es/components/code-snippet/index.js';
import '@carbon/web-components/es/components/button/index.js';
import '@carbon/web-components/es/components/icon/index.js';
import '@carbon/web-components/es/components/icon-button/index.js';
import '@carbon/web-components/es/components/icon-indicator/index.js';
import '@carbon/web-components/es/components/tag/index.js';
import '@carbon/web-components/es/components/notification/index.js';
import '@carbon/web-components/es/components/accordion/index.js';
import '@carbon/web-components/es/components/tabs/index.js';
import '@carbon/web-components/es/components/form/index.js';
import '@carbon/web-components/es/components/form-group/index.js';
import '@carbon/web-components/es/components/text-input/index.js';
import '@carbon/web-components/es/components/password-input/index.js';
import '@carbon/web-components/es/components/search/index.js';
import '@carbon/web-components/es/components/number-input/index.js';
import '@carbon/web-components/es/components/textarea/index.js';
import '@carbon/web-components/es/components/select/index.js';
import '@carbon/web-components/es/components/checkbox/index.js';
import '@carbon/web-components/es/components/dropdown/index.js';
import '@carbon/web-components/es/components/data-table/index.js';
import '@carbon/web-components/es/components/pagination/index.js';
import '@carbon/web-components/es/components/tile/index.js';

import { goWasmBuildCarbonIcons } from './icons-generated.js';

globalThis.goWasmBuildCarbonIcons = goWasmBuildCarbonIcons;
globalThis.goWasmBuildCarbonIcon = (name, size = 16) => {
    const entry = goWasmBuildCarbonIcons[name];
    if (!entry) {
        return undefined;
    }
    return entry[size] || entry[16] || entry[20] || entry[24] || entry[32];
};
