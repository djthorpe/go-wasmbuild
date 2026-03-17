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

import {
    Add16,
    Add20,
    Add24,
    Add32,
    ArrowRight16,
    ArrowRight20,
    ArrowRight24,
    ArrowRight32,
    Favorite16,
    Favorite20,
    Favorite24,
    Favorite32,
    Launch16,
    Launch20,
    Launch24,
    Launch32,
    Search16,
    Search20,
    Search24,
    Search32,
    Settings16,
    Settings20,
    Settings24,
    Settings32,
    UserAvatar16,
    UserAvatar20,
    UserAvatar24,
    UserAvatar32,
    WarningFilled16,
    WarningFilled20,
    WarningFilled24,
    WarningFilled32,
} from '@carbon/icons/es';

const goWasmBuildCarbonIcons = {
    add: { 16: Add16, 20: Add20, 24: Add24, 32: Add32 },
    'arrow-right': { 16: ArrowRight16, 20: ArrowRight20, 24: ArrowRight24, 32: ArrowRight32 },
    favorite: { 16: Favorite16, 20: Favorite20, 24: Favorite24, 32: Favorite32 },
    launch: { 16: Launch16, 20: Launch20, 24: Launch24, 32: Launch32 },
    search: { 16: Search16, 20: Search20, 24: Search24, 32: Search32 },
    settings: { 16: Settings16, 20: Settings20, 24: Settings24, 32: Settings32 },
    'user--avatar': { 16: UserAvatar16, 20: UserAvatar20, 24: UserAvatar24, 32: UserAvatar32 },
    'warning--filled': { 16: WarningFilled16, 20: WarningFilled20, 24: WarningFilled24, 32: WarningFilled32 },
};

globalThis.goWasmBuildCarbonIcons = goWasmBuildCarbonIcons;
globalThis.goWasmBuildCarbonIcon = (name, size = 16) => {
    const entry = goWasmBuildCarbonIcons[name];
    if (!entry) {
        return undefined;
    }
    return entry[size] || entry[16] || entry[20] || entry[24] || entry[32];
};
