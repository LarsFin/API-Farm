require 'json'

class Config
    attr_reader :settings

    def initialize environment

        if environment == "DEV" 
            dev_file = File.read("config.dev.json")

            settings = JSON.parse(dev_file)
        end

        # read respective config file. E.g; if DEV; should read from config file 'config.dev.json'. This file will be at the root of the lang/framework dir.
        # convert read data into hash using json library
        # set converted data to settings attribute
        # unit tests too

        # edge cases
        # if a config file doesn't exist for used environment; throw error and stop program
    end
end