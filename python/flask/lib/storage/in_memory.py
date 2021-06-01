'''TODO'''
class InMemory():
    '''in memory storage'''
    def __init__(self, video_games):
        '''init method'''
        self.video_games = video_games

    def add(self, video_game):
        '''Adds a video game object to the video games array'''
        self.video_games.append(video_game)

    def get(self, index):
        '''Gets a video game object from the video games list'''
        for self.video_game in self.video_games:
            if self.video_game.id == index:
                print("I am in get")
                return self.video_game
        