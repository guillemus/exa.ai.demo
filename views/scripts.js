/** @type {const} */
const searchTypeValues = ["instant", "fast", "auto", "deep"]

/**
 * @param {number} index
 * @returns {string}
 */
function searchTypePercentForIndex(index) {
    return ((index / (searchTypeValues.length - 1)) * 100) + "%"
}

/**
 * @param {HTMLElement} root
 * @param {string} value
 * @param {boolean} snap
 */
function setSearchTypeValue(root, value, snap) {
    const index = searchTypeValues.indexOf(value)
    if (index < 0) return

    if (snap) root.style.setProperty("--slider-position", searchTypePercentForIndex(index))

    const input = /** @type {HTMLInputElement} */ (root.querySelector("[data-search-type-value]"))
    input.value = value
    input.dispatchEvent(new Event("input", { bubbles: true }))
}

window.setSearchTypeValue = setSearchTypeValue

/**
 * @param {HTMLElement} root
 * @param {PointerEvent} event
 * @param {boolean} snap
 */
function updateSearchTypeFromPointer(root, event, snap) {
    const track = /** @type {HTMLElement} */ (root.querySelector("[data-search-type-track]"))
    const rect = track.getBoundingClientRect()
    const pct = Math.max(0, Math.min(1, (event.clientX - rect.left) / rect.width))
    const index = Math.round(pct * (searchTypeValues.length - 1))

    root.style.setProperty(
        "--slider-position",
        snap ? searchTypePercentForIndex(index) : (pct * 100) + "%",
    )
    setSearchTypeValue(root, searchTypeValues[index], false)
}

window.updateSearchTypeFromPointer = updateSearchTypeFromPointer

/**
 * @param {HTMLElement} element
 * @returns {void}
 */
function initSearchTypeSlider(element) {
    if (element.dataset.searchTypeReady === "true") return
    element.dataset.searchTypeReady = "true"

    const input = /** @type {HTMLInputElement} */ (element.querySelector("[data-search-type-value]"))
    setSearchTypeValue(element, input.value, true)

    for (const option of element.querySelectorAll("[data-search-type-option]")) {
        const optionElement = /** @type {HTMLElement} */ (option)
        optionElement.addEventListener("click", () => {
            setSearchTypeValue(element, optionElement.dataset.searchTypeOption ?? "auto", true)
        })
    }

    const track = /** @type {HTMLElement} */ (element.querySelector("[data-search-type-track]"))
    track.addEventListener("pointerdown", event => {
        track.setPointerCapture(event.pointerId)
        updateSearchTypeFromPointer(element, event, false)

        /** @param {PointerEvent} moveEvent */
        const move = moveEvent => updateSearchTypeFromPointer(element, moveEvent, false)

        /** @param {PointerEvent} upEvent */
        const up = upEvent => {
            track.releasePointerCapture(event.pointerId)
            updateSearchTypeFromPointer(element, upEvent, true)
            track.removeEventListener("pointermove", move)
            track.removeEventListener("pointerup", up)
            track.removeEventListener("pointercancel", up)
        }

        track.addEventListener("pointermove", move)
        track.addEventListener("pointerup", up)
        track.addEventListener("pointercancel", up)
    })
}

window.initSearchTypeSlider = initSearchTypeSlider
