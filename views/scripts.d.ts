export {}

declare global {
    interface Window {
        setSearchTypeValue: (root: HTMLElement, value: string, snap: boolean) => void
        copyToClipboard: (text: string) => Promise<void>
        syncSignalsToURL: (signals: Record<string, string | number | boolean>) => void
        initSearchTypeSlider: (element: HTMLElement) => void
    }
}
