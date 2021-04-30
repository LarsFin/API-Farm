# frozen_string_literal: true

# model with attributes of a video game
class VideoGame
    attr_accessor :id,
                  :name,
                  :developers,
                  :publishers,
                  :directors,
                  :producers,
                  :designers,
                  :programmers,
                  :artists,
                  :composers,
                  :platforms,
                  :date_released

    def to_hash
        date_released_s = @date_released.strftime '%d/%m/%Y'

        {
          'id' => @id, 'name' => @name,
          'developers' => @developers, 'publishers' => @publishers,
          'directors' => @directors, 'producers' => @producers,
          'designers' => @designers, 'programmers' => @programmers,
          'artists' => @artists, 'composers' => @composers,
          'platforms' => @platforms, 'date_released' => date_released_s
        }
    end
end