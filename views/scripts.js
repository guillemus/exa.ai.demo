/* global tippy */

/** @type {readonly string[]} */
const searchTypeValues = ['instant', 'fast', 'auto', 'deep']

/** @type {readonly number[]} */
const searchTypePositions = [0, 25, 50, 100]

/** @type {number | undefined} */
let searchTypeSnapFrame

/**
 * @param {number} index
 * @returns {string}
 */
function searchTypePercentForIndex(index) {
    return searchTypePositions[index] + '%'
}

/**
 * @param {HTMLElement} root
 * @param {string} value
 * @param {boolean} snap
 */
function setSearchTypeValue(root, value, snap) {
    const index = searchTypeValues.indexOf(value)
    if (index < 0) return

    if (snap) snapSearchTypeSlider(root, index)

    const input = /** @type {HTMLInputElement} */ (root.querySelector('[data-search-type-value]'))
    input.value = value
    input.dispatchEvent(new Event('input', { bubbles: true }))
}

window.setSearchTypeValue = setSearchTypeValue

/**
 * @param {HTMLElement} root
 * @param {number} index
 */
function snapSearchTypeSlider(root, index) {
    if (searchTypeSnapFrame !== undefined) window.cancelAnimationFrame(searchTypeSnapFrame)

    const start = currentSearchTypePercent(root)
    const end = searchTypePositions[index]
    const duration = 140
    const startedAt = performance.now()

    /** @param {number} now */
    const tick = (now) => {
        const progress = Math.min(1, (now - startedAt) / duration)
        const eased = 1 - Math.pow(1 - progress, 3)
        root.style.setProperty('--slider-position', start + (end - start) * eased + '%')
        if (progress < 1) {
            searchTypeSnapFrame = window.requestAnimationFrame(tick)
            return
        }
        searchTypeSnapFrame = undefined
    }

    searchTypeSnapFrame = window.requestAnimationFrame(tick)
}

/**
 * @param {HTMLElement} root
 * @returns {number}
 */
function currentSearchTypePercent(root) {
    const value =
        root.style.getPropertyValue('--slider-position') ||
        getComputedStyle(root).getPropertyValue('--slider-position')
    const parsed = Number.parseFloat(value)
    if (Number.isNaN(parsed)) return 0
    return parsed
}

/**
 * @param {number} pct
 * @returns {number}
 */
function nearestSearchTypeIndex(pct) {
    let nearestIndex = 0
    let nearestDistance = Math.abs(pct - searchTypePositions[0])
    for (let index = 1; index < searchTypePositions.length; index++) {
        const distance = Math.abs(pct - searchTypePositions[index])
        if (distance < nearestDistance) {
            nearestIndex = index
            nearestDistance = distance
        }
    }
    return nearestIndex
}

/**
 * @param {HTMLElement} root
 * @param {PointerEvent} event
 * @param {boolean} snap
 */
function updateSearchTypeFromPointer(root, event, snap) {
    const track = /** @type {HTMLElement} */ (root.querySelector('[data-search-type-track]'))
    const rect = track.getBoundingClientRect()
    const pct = Math.max(0, Math.min(100, ((event.clientX - rect.left) / rect.width) * 100))
    const index = nearestSearchTypeIndex(pct)

    if (snap) {
        snapSearchTypeSlider(root, index)
    } else {
        if (searchTypeSnapFrame !== undefined) window.cancelAnimationFrame(searchTypeSnapFrame)
        searchTypeSnapFrame = undefined
        root.style.setProperty('--slider-position', pct + '%')
    }
    setSearchTypeValue(root, searchTypeValues[index], false)
}

/** @type {WeakMap<HTMLElement, number>} */
const copyResetTimers = new WeakMap()

/**
 * @param {string} text
 * @param {HTMLElement=} button
 * @returns {Promise<void>}
 */
async function copyToClipboard(text, button) {
    await navigator.clipboard.writeText(text)
    if (!button) return

    button.classList.add('is-copied')

    const existingTimer = copyResetTimers.get(button)
    if (existingTimer !== undefined) window.clearTimeout(existingTimer)

    const timer = window.setTimeout(() => {
        button.classList.remove('is-copied')
        copyResetTimers.delete(button)
    }, 1000)
    copyResetTimers.set(button, timer)
}

window.copyToClipboard = copyToClipboard

/** @type {Record<string, string | number | boolean>} */
const urlSignalDefaults = {
    query: 'Latest news on Nvidia',
    searchType: 'auto',
    deepModel: 'deep',
    numResults: 10,
    category: 'company',
    structuredOutputs: false,
    streamResponse: false,
    systemPromptEnabled: false,
    systemPrompt: '',
    highlights: true,
    highlightMaxCharacters: 4000,
    highlightQuery: '',
    text: false,
    textMaxCharacters: 20000,
    textMainContentOnly: true,
    maxAgeHours: '',
    livecrawlTimeout: 10000,
    includeDomains: '',
    excludeDomains: '',
    startPublishedDate: '',
    endPublishedDate: '',
    userLocation: '',
}

let urlSignalsReady = false

/**
 * @param {Record<string, string | number | boolean>} signals
 * @returns {void}
 */
function syncSignalsToURL(signals) {
    if (!urlSignalsReady) {
        urlSignalsReady = true
        return
    }

    const params = new URLSearchParams()
    for (const [key, value] of Object.entries(signals)) {
        if (!(key in urlSignalDefaults)) continue
        if (String(value) === String(urlSignalDefaults[key])) continue
        if (value === '') continue
        params.set(key, String(value))
    }

    const query = params.toString()
    const url = location.pathname + (query ? '?' + query : '')
    history.replaceState(null, '', url)
}

window.syncSignalsToURL = syncSignalsToURL

/**
 * @param {HTMLElement} element
 * @returns {void}
 */
function initSearchTypeSlider(element) {
    if (element.dataset.searchTypeReady === 'true') return
    element.dataset.searchTypeReady = 'true'

    const input = /** @type {HTMLInputElement} */ (
        element.querySelector('[data-search-type-value]')
    )
    setSearchTypeValue(element, input.value, true)

    for (const option of element.querySelectorAll('[data-search-type-option]')) {
        const optionElement = /** @type {HTMLElement} */ (option)
        optionElement.addEventListener('click', () => {
            setSearchTypeValue(element, optionElement.dataset.searchTypeOption ?? 'auto', true)
        })
    }

    const track = /** @type {HTMLElement} */ (element.querySelector('[data-search-type-track]'))
    track.addEventListener('pointerdown', (event) => {
        track.setPointerCapture(event.pointerId)
        updateSearchTypeFromPointer(element, event, false)

        /** @param {PointerEvent} moveEvent */
        const move = (moveEvent) => updateSearchTypeFromPointer(element, moveEvent, false)

        /** @param {PointerEvent} upEvent */
        const up = (upEvent) => {
            track.releasePointerCapture(event.pointerId)
            updateSearchTypeFromPointer(element, upEvent, true)
            track.removeEventListener('pointermove', move)
            track.removeEventListener('pointerup', up)
            track.removeEventListener('pointercancel', up)
        }

        track.addEventListener('pointermove', move)
        track.addEventListener('pointerup', up)
        track.addEventListener('pointercancel', up)
    })
}

window.initSearchTypeSlider = initSearchTypeSlider

/**
 * @param {HTMLElement} element
 * @returns {void}
 */
function initTooltip(element) {
    if (element.dataset.tooltipReady === 'true') return
    element.dataset.tooltipReady = 'true'

    tippy(element, {
        content: element.dataset.tooltip ?? '',
        allowHTML: true,
        arrow: true,
        delay: [120, 0],
        interactive: true,
        maxWidth: 260,
        placement: 'top',
        theme: 'exa',
    })
}

window.initTooltip = initTooltip
