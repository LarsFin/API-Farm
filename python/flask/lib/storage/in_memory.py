''' in memory storage '''
class InMemory():
    def __init__(self, video_games):
        self.video_games = video_games

    def add(self, video_game):
        self.video_games.append(video_game)
        