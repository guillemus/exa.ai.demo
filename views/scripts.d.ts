export {}

declare global {
    const tippy: (element: HTMLElement, options: Record<string, unknown>) => void
    interface Window {
        setSearchTypeValue: (root: HTMLElement, value: string, snap: boolean) => void
        copyToClipboard: (text: string, button?: HTMLElement) => Promise<void>
        syncQueryParam: (name: string, value: string | number | boolean | null | undefined) => void
        initSearchTypeSlider: (element: HTMLElement) => void
        initTooltip: (element: HTMLElement) => void
    }
}
