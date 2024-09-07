export const Spotlight = (() => {
  let focusActive = false
  let focusOverlay = null
  let spotlightRadius = 100 // 聚光燈初始半徑大小
  const minRadius = 20      // 聚光燈最小半徑
  const maxRadius = 300     // 聚光燈最大半徑

  // 創建遮罩層
  function createOverlay() {
    focusOverlay = document.createElement('div')
    focusOverlay.id = "spotlight"
    focusOverlay.style.position = 'fixed'
    focusOverlay.style.top = '0'
    focusOverlay.style.left = '0'
    focusOverlay.style.width = '100%'
    focusOverlay.style.height = '100%'
    focusOverlay.style.pointerEvents = 'none'
    focusOverlay.style.zIndex = '1000'
    focusOverlay.style.background = 'rgba(0, 0, 0, 0.5)'  // 陰影部分
    document.body.appendChild(focusOverlay)
  }

  function updateFocus(event) {
    if (!focusActive) return

    // 取得滑鼠座標
    const mouseX = event.clientX
    const mouseY = event.clientY

    // 設置遮罩的聚焦區域 (亮的部分)
    focusOverlay.style.background = `
      radial-gradient(circle at ${mouseX}px ${mouseY}px,
      rgba(255, 255, 255, 0) ${spotlightRadius}px,
      rgba(0, 0, 0, 0.75) ${spotlightRadius + 50}px)
    `;
  }

  function toggleFocus(event) {
    if (event.key === 'z' || event.key === 'Z') {
      focusActive = !focusActive

      if (focusActive) {
        if (!focusOverlay) {
          createOverlay()
        }
        focusOverlay.style.display = 'block'
        document.addEventListener('mousemove', updateFocus)
        document.addEventListener('wheel', adjustSpotlightSize);  // 監聽滾輪事件
      } else {
        focusOverlay.style.display = 'none'
        document.removeEventListener('mousemove', updateFocus)
        document.removeEventListener('wheel', adjustSpotlightSize);  // 停止監聽滾輪事件
      }
    }
  }

  function adjustSpotlightSize(event) {
    if (!event.shiftKey) return;  // altKey, ctrlKey, metaKey

    if (event.deltaY < 0) {
      // 往上滾動, 放大聚光燈
      spotlightRadius = Math.min(spotlightRadius + 10, maxRadius)
    } else {
      // 往下滾動, 縮小
      spotlightRadius = Math.max(spotlightRadius - 10, minRadius)
    }

    // 更新聚光燈大小
    updateFocus(event)
  }

  return {
    id: 'spotlight',
    init: (reveal) => {
      document.addEventListener('keydown', toggleFocus)
    }
  }
})()
