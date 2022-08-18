require 'Assets.Player.player_vars'

local movement = {}

function movement.animation(self)
  local anim
  
  local updateAnimation_fun_t = {
    [self.walking] = function ()
      if self.dir.x > 0 then
        anim = walk_right

      elseif self.dir.x < 0 then
        anim = walk_left

      elseif self.dir.z > 0 then
        anim = walk_down

      elseif self.dir.z < 0 then
        anim = walk_up

      else
        if self.last_dir == dir_down then
          anim = idle_down

        elseif self.last_dir == dir_left then
          anim = idle_left
          
        elseif self.last_dir == dir_right then
          anim = idle_right

        elseif self.last_dir == dir_up then
          anim = idle_up
        end
      end

      return anim
    end,
  }

  if updateAnimation_fun_t[self.state] then
    return updateAnimation_fun_t[self.state]()
  else
    print('no move function for: ', self.state)
  end
end

function movement.move(self, dt, old_x, old_z)
  local x, z, dir

  self.speed = self.regular_speed
  dir = self.dir

  if dir ~= vmath.vector3() then
    dir = vmath.normalize(dir)
  end

  x = old_x + self.dx * self.speed * dt
  z = old_z + self.dz * self.speed * dt

  self.dx = self.dx * (1 - self.friction)
  self.dz = self.dz * (1 - self.friction)

  local kx = dir.x
  local kz = dir.z

  self.dx = self.dx + kx
  self.dz = self.dz + kz

  return x, z
end

return movement