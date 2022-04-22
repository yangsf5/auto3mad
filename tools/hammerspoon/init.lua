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

local function keyBind(hyper, keyFnTable)
	for key,fn in pairs(keyFnTable) do
		hs.hotkey.bind(hyper, key, fn)
	end
end

keyBind({}, {
	f1 = function() module.moveWindowToMonitor("MI") end,
  f2 = function() module.moveWindowToMonitor("DELL") end,
	f3 = function() module.moveWindowToMonitor("Retina") end,
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


hs.alert.show("Hammerspoon config loaded")