-- Open the store page for MC
process.exec("cmd /C start ms-windows-store://pdp/?PFN=Microsoft.MinecraftUWP_8wekyb3d8bbwe")

-- Wait some time until the poster shows up
screen.waitUntilVisible("ms_store_mc10_poster.png")

-- There is a animation for the poster scroll
util.sleep(1500)

-- Check if the owned mark is there
if screen.isVisible("ms_store_owned.png") then
    screen.clickOn("ms_store_button.png")

    log.info("Downloading minecraft...")

    screen.waitUntilVisible("ms_store_downloaded.png")
    screen.waitUntilVisible("ms_store_button.png", "ms_store_after_dl.png")
end

-- Close the store
local pid = process.waitFor("WinStore.App.exe")
process.kill(pid)
