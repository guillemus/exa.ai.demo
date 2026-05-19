export {}

declare global {
    interface Window {
        setSearchTypeValue: (root: HTMLElement, value: string, snap: boolean) => void
        updateSearchTypeFromPointer: (root: HTMLElement, event: PointerEvent, snap: boolean) => void
        initSearchTypeSlider: (element: HTMLElement) => void
    }
}
