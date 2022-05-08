hs.window.animationDuration = 0
hs.grid.setMargins("0, 0")


module = {}

module.maximizeWindow = function()
	hs.grid.maximizeWindow(hs.window.focusedWindow())
end

module.centerOnScreen = function()
	local window = hs.window.focusedWindow()
	window:centerOnScreen(window:screen())
end

module.moveWindowToMonitor = function(monitor)
	local screen = hs.screen.find(monitor)
	hs.window.focusedWindow():moveToScreen(screen)
  module.maximizeWindow()
end

module.moveMouseToMonitor = function(monitor)
	local screen = hs.screen.find(monitor)
	local rect = screen:fullFrame()
	local center = hs.geometry.rectMidPoint(rect)

	hs.mouse.absolutePosition(center)
	hs.window.highlight.start()
end

module.leftHalf = function()
	local window = hs.window.focusedWindow()
	local screen = window:screen()
	local grid = hs.grid.getGrid(screen)
	local cell = hs.geometry(0, 0, 0.5 * grid.w, grid.h)
	hs.grid.set(window, cell, screen)
end

module.rightHalf = function()
	local window = hs.window.focusedWindow()
	local screen = window:screen()
	local grid = hs.grid.getGrid(screen)
	local cell = hs.geometry(0.5 * grid.w, 0, 0.5 * grid.w, grid.h)
	hs.grid.set(window, cell, screen)
end

-- 坐标原点 (0,0) 为屏幕左上角
-- (x, y) 为窗口的左上角坐标，值为比率
-- w, h 为宽高的尺寸，值为比率
module.moveWindowInMonitor = function(x, y, w, h)
	local window = hs.window.focusedWindow()
	local screen = window:screen()
	local grid = hs.grid.getGrid(screen)
	local cell = hs.geometry(grid.w * x, grid.h * y, grid.w * w, grid.h * h)
	hs.grid.set(window, cell, screen)
end

-- 编码时的窗口布局
function coding()
  app = hs.application.find("Code") -- Visual Studio Code
  wins = app:allWindows()
  hs.fnutils.each(wins, function(win) 
    if not win:isMinimized() then
      win:moveToScreen(hs.screen.find("dell")):maximize():focus() 
    end
  end)
end

-- 写作时的窗口布局
function writing()
  app = hs.application.find("语雀")

  goalWin = app:findWindow("三层目标")
  if goalWin == nil then
    -- TODO 看看怎么自动打开这个窗口
    hs.alert.show("请打开语雀「三层目标」")
    return
  end

  goalMonitor = hs.screen.find("mi")
  if goalMonitor == nil then
    goalMonitor = hs.screen.find("dell")
  end
  if goalMonitor == nil then
    hs.alert.show("请连接小米或戴尔显示器")
  end

  goalWin:moveToScreen(goalMonitor):maximize():focus()

  app:findWindow("语雀"):moveToScreen(hs.screen.find("retina")):maximize():focus()
end

local function keyBind(hyper, keyFnTable)
	for key,fn in pairs(keyFnTable) do
		hs.hotkey.bind(hyper, key, fn)
	end
end

keyBind({}, {
	f1 = function() module.moveWindowToMonitor("MI") end,
  f2 = function() module.moveWindowToMonitor("DELL") end,
	f3 = function() module.moveWindowToMonitor("Retina") end,
  f4 = coding,
  f5 = writing,
  f10 = function() hs.audiodevice.defaultOutputDevice():setOutputMuted(true) end,
})


keyBind(nil, {
  home = function() module.moveMouseToMonitor("MI") end, -- fn+向左箭头
	pageup = function() module.moveMouseToMonitor("DELL") end, -- fn+向上箭头
  pagedown = function() module.moveMouseToMonitor("Retina") end, -- fn+向下箭头
})

keyBind({"rightalt"}, {
	m = module.maximizeWindow,
	down = module.centerOnScreen,
	left = module.leftHalf,
	right = module.rightHalf,
})

keyBind({"rightshift"}, {
	left = function() module.moveWindowInMonitor(0, 0, 0.33, 1) end,
  down = function() module.moveWindowInMonitor(0.33, 0, 0.33, 1) end,
	right = function() module.moveWindowInMonitor(0.66, 0, 0.33, 1) end,
})


-- 监控 WIFI 变化
function wifiChanged(watcher, message, interface)
  wName = hs.wifi.currentNetwork()
  print(message, " ", wName)
end

w = hs.wifi.watcher.new(wifiChanged)
--w:watchingFor("all")
w:start()


-- 监控显示器拔插、分辨率等变化
function monitorChanged()
  all = hs.screen.allScreens()
  hs.fnutils.each(all, function(v) print(v:name()) end)
end

hs.screen.watcher.new(monitorChanged):start()


hs.alert.show("Hammerspoon config loaded")