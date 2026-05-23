export {}

declare global {
    const tippy: (element: HTMLElement, options: Record<string, unknown>) => void
    interface Window {
        setSearchTypeValue: (root: HTMLElement, value: string, snap: boolean) => void
        copyToClipboard: (text: string, button?: HTMLElement) => Promise<void>
        syncSignalsToURL: (signals: Record<string, string | number | boolean>) => void
        initSearchTypeSlider: (element: HTMLElement) => void
        initTooltip: (element: HTMLElement) => void
    }
}
