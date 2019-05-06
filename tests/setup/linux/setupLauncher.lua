-- Function to check and setup a profile
function checkAndSetupProfile()
    -- Open the profile editor
    screen.clickOn("launcher_profile_edit.png")

    -- Wait a bit until the dropdown pops up
    screen.waitUntilVisible("launcher_profile_version_dropdown.png")
    screen.clickOn("launcher_profile_version_dropdown.png")
    util.sleep(500)
    screen.debug()

    -- Wait a second and take a debug screen
    local x,y = screen.waitUntilVisible("launcher_profile_latest_version.png", "launcher_profile_latest_version_other.png")
    mouse.move(x,y)

    -- Move mouse a bit outwards
    mouse.moveRelative(-14, -10)

    -- Detect latest stable
    while true do
        mouse.moveRelative(0,32)
        local version = screen.readText(250, 32)
        local s = string.find(version, "%(beta%)")
        if s == nil then
            mouse.moveRelative(10,10)
            mouse.click()
            break
        end
    end

    -- Save the profile
    screen.waitUntilVisible("launcher_profile_save.png")
    screen.clickOn("launcher_profile_save.png")
end

function startMCPE()
    -- Wait for download and play button
    screen.waitUntilVisible("launcher_download_and_play.png")
    screen.clickOn("launcher_download_and_play.png")
end

-- Abort the screensaver
mouse.moveRelative(10,10)
util.sleep(500)
mouse.moveRelative(-10,-10)

-- Start the launcher
process.run("mcpelauncher-ui-qt")

-- Wait until the launcher pops up
util.sleep(2000)

-- Configure XFCE
if screen.isVisible("xfce_setup.png") then
    screen.clickOn("xfce_setup.png")
end

-- Check if profile is there
if screen.isVisible("launcher_profile_edit.png") then
    -- Ensure we have a proper profile setup
    checkAndSetupProfile()

    -- Start the game
    startMCPE()
    return
end

-- Launcher signin is visible?
if screen.isVisible("launcher_signin.png") == false then
    screen.debug()
end

-- Wait until the signin button arrives
screen.waitUntilVisible("launcher_signin.png")

-- Click on the button
screen.clickOn("launcher_signin.png")
mouse.moveRelative(0, 150)
mouse.click()
screen.debug()

-- Wait until email field comes up
screen.waitUntilVisible("google_email_field.png")
screen.clickOn("google_email_field.png")

-- Account name
keyboard.type("__YOUR_EMAIL__", true)
util.sleep(666)
keyboard.press("enter")

util.sleep(2000)

-- Put in the password
keyboard.type("__YOUR_PASSWORD__", true)
util.sleep(666)
keyboard.press("enter")

-- Now shit hits the desk, the launcher doesn't login
util.sleep(500)

-- Kill it with fire
local pid = process.waitFor("mcpelauncher-ui-qt")
process.kill(pid)

-- Start the launcher
process.run("mcpelauncher-ui-qt")

-- Wait for it to come back up
util.sleep(2000)

-- Agree TOS
screen.waitUntilVisible("google_tos_agree.png")
screen.clickOn("google_tos_agree.png")

-- Wait until profile selector pops up
screen.waitUntilVisible("launcher_profile_edit.png")

-- Ensure we have a correctly setup profile
checkAndSetupProfile()

-- Start
startMCPE()