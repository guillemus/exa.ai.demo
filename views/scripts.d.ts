export {}

declare global {
    interface Window {
        setSearchTypeValue: (root: HTMLElement, value: string, snap: boolean) => void
        copyToClipboard: (text: string) => Promise<void>
        initSearchTypeSlider: (element: HTMLElement) => void
    }
}
