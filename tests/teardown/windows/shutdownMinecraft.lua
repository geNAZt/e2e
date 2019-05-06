-- Kill MC
local pid = process.waitFor("Minecraft.Windows.exe")
process.kill(pid)