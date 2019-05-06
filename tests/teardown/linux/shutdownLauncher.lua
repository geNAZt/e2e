-- Kill MC launcher
local pid = process.waitFor("mcpelauncher-ui-qt")
process.kill(pid)