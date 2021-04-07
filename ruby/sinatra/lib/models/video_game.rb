# frozen_string_literal: true

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
        {
            'id' => @id,
            'name' => @name,
            'developers' => @developers,
            'publishers' => @publishers,
            'directors' => @directors,
            'producers' => @producers,
            'designers' => @designers,
            'programmers' => @programmers,
            'artists' => @artists,
            'composers' => @composers,
            'platforms' => @platforms,
            'date_released' => @date_released
        }
    end
end