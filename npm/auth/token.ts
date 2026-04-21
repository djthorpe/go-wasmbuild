import type { TokenResponse } from "./auth";

const DEFAULT_TOKEN_KEY = "auth:token";
const DEFAULT_EXPIRY_SKEW_MS = 30_000;

export interface StoredTokenResponse extends TokenResponse {
    stored_at?: string;
}

export interface TokenOptions {
    key?: string;
    storage?: Storage;
    expirySkewMs?: number;
}

export default class Token {
    readonly key: string;
    readonly storage?: Storage;
    readonly expirySkewMs: number;

    constructor(options: TokenOptions = {}) {
        this.key = options.key ?? DEFAULT_TOKEN_KEY;
        this.storage = options.storage ?? getDefaultStorage();
        this.expirySkewMs = options.expirySkewMs ?? DEFAULT_EXPIRY_SKEW_MS;
    }

    read(): StoredTokenResponse | undefined {
        const raw = this.storage?.getItem(this.key);
        if (!raw) {
            return undefined;
        }
        try {
            const value = JSON.parse(raw) as StoredTokenResponse;
            if (!value || typeof value !== "object" || typeof value.access_token !== "string") {
                this.delete();
                return undefined;
            }
            return value;
        } catch {
            this.delete();
            return undefined;
        }
    }

    write(token: TokenResponse): StoredTokenResponse {
        const storedToken: StoredTokenResponse = {
            ...token,
            stored_at: new Date().toISOString(),
        };
        this.storage?.setItem(this.key, JSON.stringify(storedToken));
        return storedToken;
    }

    delete(): void {
        this.storage?.removeItem(this.key);
    }

    valid(token: StoredTokenResponse | undefined = this.read()): boolean {
        if (!token?.access_token) {
            return false;
        }

        const now = Date.now() + this.expirySkewMs;

        if (token.expiry) {
            const expiry = Date.parse(token.expiry);
            if (Number.isFinite(expiry)) {
                return expiry > now;
            }
        }

        if (typeof token.expires_in === "number" && token.stored_at) {
            const storedAt = Date.parse(token.stored_at);
            if (Number.isFinite(storedAt)) {
                return storedAt + token.expires_in * 1000 > now;
            }
        }

        return true;
    }
}

function getDefaultStorage(): Storage | undefined {
    try {
        return globalThis.sessionStorage;
    } catch {
        return undefined;
    }
}