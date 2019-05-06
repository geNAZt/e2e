-- Start minecraft
process.exec("cmd /C start minecraft:")

-- Wait until the mojang screen pops up
screen.waitUntilVisible("mc_start_logo.png")
screen.waitUntilNotVisible("mc_start_logo.png")
util.sleep(2000)

-- Check if we need to login into xbox
if screen.isVisible("mc_xbox_button.png") then
    log.info("Logging in into XBOX Live...")
    screen.clickOn("mc_xbox_button.png")
end

-- Wait until the community icon comes online