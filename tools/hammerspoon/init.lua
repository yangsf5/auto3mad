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

  app = hs.application.find("iTerm")
  wins = app:allWindows()
  hs.fnutils.each(wins, function(win) 
    if not win:isMinimized() then
      win:moveToScreen(hs.screen.find("mi")):maximize():focus()
    end
  end)
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
  f10 = function() hs.audiodevice.defaultOutputDevice():setOutputMuted(true) end,
})

keyBind({"rightalt"}, {
	m = module.maximizeWindow,
	-- down = module.centerOnScreen, -- 会干扰 XMind 向下移动节点，所以取消这个快捷键
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