--#!/usr/local/bin/lua
--[[
local num = redis.call('GET', KEYS[1]);
if not num then
    redis.call('SET',KEYS[1], 1);
    return 1;
else
    local res = num + 1;
    redis.call('SET',KEYS[1], res);
    return res;
end
]]--
--[[
function stat(dur_data, now)
    if not dur_data then
        dur_data = {}
        dur_data["updated_at"] = now
    else
        local diff = now - dur_data["updated_at"]
        if 1000000000 < diff and diff < 10000000000 then
            if not dur_data["duration"] then
                dur_data["duration"] = 0
            end
            dur_data["duration"] = dur_data["duration"] + diff
        end
        dur_data["updated_at"] = now
    end
    print("dur_data.updated_at=", dur_data["updated_at"])
    print("dur_data.duration=", dur_data["duration"])
end

print("test 1:")
stat(nil, 70000000000)
print("test 2:")
local dur_data = {updated_at=65000000000}
stat(dur_data, 70000000000)
print("test 3:")
dur_data = {updated_at=65000000000,duration=5000000000}
stat(dur_data, 70000000000)
]]--

--[[
function serialize(obj)
    local lua = ""
    local t = type(obj)
    if t == "number" then
        lua = lua .. obj
    elseif t == "boolean" then
        lua = lua .. tostring(obj)
    elseif t == "string" then
        lua = lua .. string.format("%q", obj)
    elseif t == "table" then
        lua = lua .. "{\n"
        for k, v in pairs(obj) do
            lua = lua .. "[" .. serialize(k) .. "]=" .. serialize(v) .. ",\n"
        end
        local metatable = getmetatable(obj)
        if metatable ~= nil and type(metatable.__index) == "table" then
            for k, v in pairs(metatable.__index) do
                lua = lua .. "[" .. serialize(k) .. "]=" .. serialize(v) .. ",\n"
            end
        end
        lua = lua .. "}"
    elseif t == "nil" then
        return nil
    else
        error("can not serialize a " .. t .. " type.")
    end
    return lua
end

function unserialize(lua)
    local t = type(lua)
    if t == "nil" or lua == "" then
        return nil
    elseif t == "number" or t == "string" or t == "boolean" then
        lua = tostring(lua)
    else
        error("can not unserialize a " .. t .. " type.")
    end
    lua = "return " .. lua
    --local func = loadstring(lua) loadstring在lua5.2中已经被弃用了
    local func = load(lua)
    if func == nil then
        return nil
    end
    return func()
end

-- 测试代码如下
data = {["a"] = "a", ["b"] = "b", [1] = 1, [2] = 2, ["t"] = {1, 2, 3}}
local sz = serialize(data)
print(sz)
print("---------")
print(serialize(unserialize(sz)))
]]--

--[[
--local mt = setmetatable(_G, nil)
local function serialize(obj)
    local lua = ""
    local t = type(obj)
    if t == "number" then
        lua = lua .. obj
    elseif t == "boolean" then
        lua = lua .. tostring(obj)
    elseif t == "string" then
        lua = lua .. string.format("%q", obj)
    elseif t == "table" then
        lua = lua .. "{\n"
        for k, v in pairs(obj) do
            lua = lua .. "[" .. serialize(k) .. "]=" .. serialize(v) .. ",\n"
        end
        local metatable = getmetatable(obj)
        if metatable ~= nil and type(metatable.__index) == "table" then
            for k, v in pairs(metatable.__index) do
                lua = lua .. "[" .. serialize(k) .. "]=" .. serialize(v) .. ",\n"
            end
        end
        lua = lua .. "}"
    elseif t == "nil" then
        return nil
    else
        error("can not serialize a " .. t .. " type.")
    end
    return lua
end
local function unserialize(lua)
    local t = type(lua)
    if t == "nil" or lua == "" then
        return nil
    elseif t == "number" or t == "string" or t == "boolean" then
        lua = tostring(lua)
    else
        error("can not unserialize a " .. t .. " type.")
    end
    lua = "return " .. lua
    local func = loadstring(lua)
    if func == nil then
        return nil
    end
    return func()
end
--setmetatable(_G, mt)
]]--
--[[
local dur_data_str = redis.call('GET', KEYS[1]);
local dur_data = unserialize(dur_data_str)
if not dur_data then
    dur_data = {}
    dur_data["updated_at"] = ARGV[1]
else
    local diff = ARGV[1] - dur_data["updated_at"]
    if 1 < diff and diff < 120 then
        if not dur_data["duration"] then
            dur_data["duration"] = 0
        end
        dur_data["duration"] = dur_data["duration"] + diff
    end
    dur_data["updated_at"] = ARGV[1]
end;
dur_data_str = serialize(dur_data)
redis.call('SET', KEYS[1], dur_data_str);
redis.call('EXPIRE', KEYS[1], 7*24*3600);
return 1
-- redis-cli  --eval test_lua.lua  "key_sunbin" , 6
]]--

--[[
local dur_data_str = redis.call('GET', KEYS[1]);
local dur_data = unserialize(dur_data_str)
if dur_data and dur_data["duration"] then
    return dur_data["duration"]
else
    return 0
end;
]]--

-- redis-cli  --eval test_lua.lua  "key_sunbin"


--local cjson = require "cjson"
--[[
local dur_data = redis.call('GET', KEYS[1])
local dur_data_str = ""
if not dur_data then
    dur_data = {
        ["updated_at"] = ARGV[1]
    }
else
    dur_data = cjson.decode(dur_data)
    local diff = ARGV[1] - dur_data["updated_at"]
    if 1 < diff and diff < 120 then
        if not dur_data["duration"] then
            dur_data["duration"] = 0
        end
        dur_data["duration"] = dur_data["duration"] + diff
    end
    dur_data["updated_at"] = ARGV[1]
end
dur_data_str = cjson.encode(dur_data)
redis.call('SET', KEYS[1], dur_data_str)
redis.call('EXPIRE', KEYS[1], 7*24*3600)
return dur_data_str
]]--

--local dur_data = redis.call('GET', KEYS[1])
--if dur_data then
--    dur_data = cjson.decode(dur_data)
--    if dur_data["duration"] then
--        return dur_data["duration"]
--    else
--        return 0
--    end
--else
--    return 0
--end

local a = "123"
local b = a - "23"
print(b)
