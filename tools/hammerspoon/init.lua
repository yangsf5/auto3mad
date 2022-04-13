hs.window.animationDuration = 0
hs.grid.setMargins("0, 0")


module = {}

module.maximizeWindow = function ()
	hs.grid.maximizeWindow(hs.window.focusedWindow())
end

module.centerOnScreen = function ()
	local window = hs.window.focusedWindow()
	window:centerOnScreen(window:screen())
end

module.moveWindowToMonitor = function(monitor)
	local screen = hs.screen.find(monitor)
	hs.window.focusedWindow():moveToScreen(screen)
end

module.moveMouseToMonitor = function(monitor)
	local screen = hs.screen.find(monitor)
	local rect = screen:fullFrame()
	local center = hs.geometry.rectMidPoint(rect)

	hs.mouse.absolutePosition(center)
	hs.window.highlight.start()
end

module.leftHalf = function ()
	local window = hs.window.focusedWindow()
	local screen = window:screen()
	local grid = hs.grid.getGrid(screen)
	local cell = hs.geometry(0, 0, 0.5 * grid.w, grid.h)
	hs.grid.set(window, cell, screen)
end

module.rightHalf = function ()
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

keyBind({"ctrl", "alt"}, {
	up = function() module.moveWindowToMonitor("DELL") end, -- ctrl+option+向上箭头
	left = function() module.moveWindowToMonitor("HP") end, -- ctrl+option+向左箭头
	down = function() module.moveWindowToMonitor("Retina") end, -- ctrl+option+向下箭头
})

keyBind(nil, {
	pageup = function() module.moveMouseToMonitor("DELL") end, -- fn+向上箭头
	home = function() module.moveMouseToMonitor("HP") end, -- fn+向左箭头
	pagedown = function() module.moveMouseToMonitor("Retina") end, -- fn+向下箭头
})

keyBind({"ctrl", "alt", "cmd"}, {
	m = module.maximizeWindow, -- ctrl+option+command+m
	c = module.centerOnScreen, -- ctrl+option+command+c
	left = module.leftHalf, -- ctrl+option+command+向左箭头
	right = module.rightHalf, -- ctrl+option+command+向右箭头
})


hs.alert.show("Hammerspoon config loaded")
