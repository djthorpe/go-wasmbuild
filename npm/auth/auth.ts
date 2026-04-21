export interface AuthOptions {
    baseUrl?: string | URL;
}

export interface AuthorizationOptions {
    state?: string;
    scope?: string;
    nonce?: string;
    loginHint?: string;
    codeVerifier?: string;
    codeChallengeMethod?: "S256" | "plain";
}

export interface AuthorizationRequest {
    url: URL;
    state: string;
    nonce?: string;
    codeVerifier: string;
    codeChallenge: string;
    codeChallengeMethod: "S256" | "plain";
}

export interface ExchangeOptions {
    nonce?: string;
    meta?: Record<string, unknown>;
}

export interface TokenResponse {
    access_token: string;
    refresh_token: string;
    token_type: string;
    expiry?: string;
    expires_in?: number;
}

export interface UserInfo {
    sub: string;
    email?: string;
    name?: string;
    groups?: string[];
    scopes?: string[];
}

export interface PublicClientConfiguration {
    issuer: string;
    client_id?: string;
}

export type PublicClientConfigurations = Record<string, PublicClientConfiguration>;

const DEFAULT_CODE_CHALLENGE_METHOD = "S256";

export class Auth {
    readonly baseUrl: URL;

    constructor(options: AuthOptions = {}) {
        const baseUrl = options.baseUrl ?? "/";
        const resolved = baseUrl instanceof URL
            ? new URL(baseUrl.toString())
            : new URL(baseUrl, globalThis.location?.origin ?? "http://localhost");
        if (!resolved.pathname.endsWith("/")) {
            resolved.pathname = `${resolved.pathname}/`;
        }
        this.baseUrl = resolved;
    }

    url(path: string, params?: Record<string, string | number | boolean | undefined>): URL {
        const normalizedPath = path.replace(/^\/+/, "");
        const url = new URL(normalizedPath, this.baseUrl);
        if (params) {
            for (const [key, value] of Object.entries(params)) {
                if (value !== undefined) {
                    url.searchParams.set(key, String(value));
                }
            }
        }
        return url;
    }

    async authorize(provider: string, redirectUri: string, options: AuthorizationOptions = {}): Promise<AuthorizationRequest> {
        const state = options.state?.trim() || randomToken(32);
        const codeVerifier = options.codeVerifier?.trim() || randomToken(48);
        const codeChallengeMethod = options.codeChallengeMethod ?? DEFAULT_CODE_CHALLENGE_METHOD;
        const codeChallenge = await createCodeChallenge(codeVerifier, codeChallengeMethod);
        const url = this.url("auth/authorize", {
            provider,
            redirect_uri: redirectUri,
            state,
            scope: options.scope,
            nonce: options.nonce,
            login_hint: options.loginHint,
            code_challenge: codeChallenge,
            code_challenge_method: codeChallengeMethod,
        });

        return {
            url,
            state,
            nonce: options.nonce,
            codeVerifier,
            codeChallenge,
            codeChallengeMethod,
        };
    }

    async config(signal?: AbortSignal): Promise<PublicClientConfigurations> {
        const response = await fetch(this.url("config"), {
            headers: {
                Accept: "application/json",
            },
            signal,
        });
        if (!response.ok) {
            throw new Error(`auth config request failed with status ${response.status}`);
        }
        return response.json() as Promise<PublicClientConfigurations>;
    }

    async exchange(
        provider: string,
        code: string,
        redirectUri: string,
        codeVerifier: string,
        options: ExchangeOptions = {},
        signal?: AbortSignal,
    ): Promise<TokenResponse> {
        const body = new URLSearchParams({
            grant_type: "authorization_code",
            provider,
            code,
            redirect_uri: redirectUri,
            code_verifier: codeVerifier,
        });
        if (options.nonce) {
            body.set("nonce", options.nonce);
        }
        if (options.meta) {
            body.set("meta", JSON.stringify(options.meta));
        }
        const response = await fetch(this.url("auth/code"), {
            method: "POST",
            headers: {
                Accept: "application/json",
            },
            body,
            signal,
        });
        if (!response.ok) {
            throw new Error(`auth exchange request failed with status ${response.status}`);
        }
        return response.json() as Promise<TokenResponse>;
    }

    async refresh(refreshToken: string, signal?: AbortSignal): Promise<TokenResponse> {
        const body = new URLSearchParams({
            grant_type: "refresh_token",
            refresh_token: refreshToken,
        });
        const response = await fetch(this.url("auth/code"), {
            method: "POST",
            headers: {
                Accept: "application/json",
            },
            body,
            signal,
        });
        if (!response.ok) {
            throw new Error(`auth refresh request failed with status ${response.status}`);
        }
        return response.json() as Promise<TokenResponse>;
    }

    async userinfo(accessToken: string, signal?: AbortSignal): Promise<UserInfo> {
        const response = await fetch(this.url("auth/userinfo"), {
            headers: {
                Accept: "application/json",
                Authorization: `Bearer ${accessToken}`,
            },
            signal,
        });
        if (!response.ok) {
            throw new Error(`auth userinfo request failed with status ${response.status}`);
        }
        return response.json() as Promise<UserInfo>;
    }

    async revoke(token: string, signal?: AbortSignal): Promise<void> {
        const body = new URLSearchParams({
            token,
        });
        const response = await fetch(this.url("auth/revoke"), {
            method: "POST",
            headers: {
                Accept: "application/json",
            },
            body,
            signal,
        });
        if (!response.ok) {
            throw new Error(`auth revoke request failed with status ${response.status}`);
        }
    }
}

function randomToken(size: number): string {
    if (size <= 0) {
        throw new Error("token size must be greater than zero");
    }
    const bytes = new Uint8Array(size);
    globalThis.crypto.getRandomValues(bytes);
    return base64UrlEncode(bytes);
}

async function createCodeChallenge(codeVerifier: string, method: "S256" | "plain"): Promise<string> {
    if (method === "plain") {
        return codeVerifier;
    }
    const data = new TextEncoder().encode(codeVerifier);
    const digest = await globalThis.crypto.subtle.digest("SHA-256", data);
    return base64UrlEncode(new Uint8Array(digest));
}

function base64UrlEncode(bytes: Uint8Array): string {
    let binary = "";
    for (const byte of bytes) {
        binary += String.fromCharCode(byte);
    }
    return globalThis.btoa(binary)
        .replace(/\+/g, "-")
        .replace(/\//g, "_")
        .replace(/=+$/g, "");
}

export default Auth;