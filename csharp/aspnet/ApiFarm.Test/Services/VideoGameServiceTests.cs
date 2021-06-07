using System;
using System.Collections.Generic;
using ApiFarm.Models.Impl;
using ApiFarm.Repositories;
using ApiFarm.Services.Impl;
using ApiFarm.Utils;
using Moq;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Services
{
    public class VideoGameServiceTests
    {
        private VideoGameService subject;

        private Mock<IRepository<VideoGame>> mockVideoGameStorage;
        private Mock<IQueryFactory> mockQueryFactory;
        private Mock<Action<VideoGame>> mockPrepVideoGame;

        [SetUp]
        protected void SetUp()
        {
            mockVideoGameStorage = new Mock<IRepository<VideoGame>>();
            mockQueryFactory = new Mock<IQueryFactory>();
            mockPrepVideoGame = new Mock<Action<VideoGame>>();

            subject = new VideoGameService(
                mockVideoGameStorage.Object,
                mockQueryFactory.Object,
                mockPrepVideoGame.Object);
        }

        private class GetShould : VideoGameServiceTests
        {
            [Test]
            public void RetrieveVideoGameAndReturnQuery()
            {
                // Arrange
                var id = 6u;
                var storedVideoGame = new VideoGame();
                var expected = new Mock<IQuery<VideoGame>>();

                mockVideoGameStorage.Setup(m => m.Get(id)).Returns(storedVideoGame);
                mockQueryFactory.Setup(m => m.Build(default, default, storedVideoGame)).Returns(expected.Object);

                // Act
                var actual = subject.Get(id);

                // Assert
                actual.ShouldBe(expected.Object);
                mockPrepVideoGame.Verify(m => m.Invoke(storedVideoGame));
            }

            [Test]
            public void ReturnUnsuccessfulQueryWhenVideoGameNotFound()
            {
                // Arrange
                var id = 99u;
                var expected = new Mock<IQuery<VideoGame>>();

                mockQueryFactory.Setup(m => m.Build<VideoGame>(404, ResponseMessages.VideoGame.NotFound(id), default)).Returns(expected.Object);

                // Act
                var actual = subject.Get(id);

                // Assert
                actual.ShouldBe(expected.Object);
            }
        }

        private class GetAllShould : VideoGameServiceTests
        {
            private readonly VideoGame storedVideoGame1 = new VideoGame();
            private readonly VideoGame storedVideoGame2 = new VideoGame();
            private readonly VideoGame storedVideoGame3 = new VideoGame();

            [Test]
            public void RetrieveVideoGamesFromStorageAndReturnQuery()
            {
                // Arrange
                var mockStoredVideoGames = new Mock<IEnumerable<VideoGame>>();
                var expected = new Mock<IQuery<IEnumerable<VideoGame>>>();

                mockStoredVideoGames.Setup(m => m.GetEnumerator()).Returns(StoredVideoGames());
                mockVideoGameStorage.Setup(m => m.GetAll()).Returns(mockStoredVideoGames.Object);
                mockQueryFactory.Setup(m => m.Build(default, default, mockStoredVideoGames.Object))
                    .Returns(expected.Object);

                // Act
                var actual = subject.GetAll();

                // Assert
                actual.ShouldBe(expected.Object);

                mockPrepVideoGame.Verify(m => m.Invoke(storedVideoGame1));
                mockPrepVideoGame.Verify(m => m.Invoke(storedVideoGame2));
                mockPrepVideoGame.Verify(m => m.Invoke(storedVideoGame3));
            }

            private IEnumerator<VideoGame> StoredVideoGames()
            {
                yield return storedVideoGame1;
                yield return storedVideoGame2;
                yield return storedVideoGame3;
            }
        }

        private class AddShould : VideoGameServiceTests
        {
            [Test]
            public void AddVideoGameToStorageAndReturnQuery()
            {
                // Arrange
                var videoGame = new VideoGame
                {
                    Name = "Vikings at Sea IV",
                    DateReleased = DateTime.Now,
                };
                var expectedQuery = new Mock<IQuery<VideoGame>>();

                mockVideoGameStorage.Setup(m => m.Add(videoGame)).Returns(videoGame);
                mockQueryFactory.Setup(m => m.Build(default, default, videoGame)).Returns(expectedQuery.Object);

                // Act
                var actual = subject.Add(videoGame);

                // Assert
                actual.ShouldBe(expectedQuery.Object);
                mockVideoGameStorage.Verify(m => m.Add(videoGame));
                mockPrepVideoGame.Verify(m => m.Invoke(videoGame));
            }

            [Test]
            public void ReturnUnsuccessfulQueryWhenVideoGameHasNoName()
            {
                // Arrange
                var videoGame = new VideoGame
                {
                    DateReleased = DateTime.Now,
                };
                var expectedQuery = new Mock<IQuery<VideoGame>>();

                mockQueryFactory.Setup(m => m.Build(400, ResponseMessages.VideoGame.RequiresName, default(VideoGame))).Returns(expectedQuery.Object);

                // Act
                var actual = subject.Add(videoGame);

                // Assert
                actual.ShouldBe(expectedQuery.Object);
            }

            [Test]
            public void ReturnUnsuccessfulQueryWhenVideoGameHasNoDateReleased()
            {
                // Arrange
                var videoGame = new VideoGame
                {
                    Name = "Vikings at Sea IV",
                };
                var expectedQuery = new Mock<IQuery<VideoGame>>();

                mockQueryFactory.Setup(m => m.Build(400, ResponseMessages.VideoGame.RequiresDateReleased, default(VideoGame))).Returns(expectedQuery.Object);

                // Act
                var actual = subject.Add(videoGame);

                // Assert
                actual.ShouldBe(expectedQuery.Object);
            }
        }

        private class UpdateShould : VideoGameServiceTests
        {
            [Test]
            public void UpdateVideoGameInStorageWithAllSetFields()
            {
                // Arrange
                var id = 5u;
                var videoGameToUpdate = new VideoGame
                {
                    Id = id,
                    Name = "Vikings at Sea IV",
                    Developers = new List<string> { "A", "B" },
                    Publishers = new List<string> { "A", "B", "C" },
                    Directors = new List<string> { "A" },
                    Producers = new List<string> { "A" },
                    Designers = new List<string> { "A", "B" },
                    Programmers = new List<string> { "A", "B", "C", "D" },
                    Artists = new List<string> { "A", "B" },
                    Composers = new List<string> { "A", "B", "C" },
                    Platforms = new List<string> { "A", "B", "C", "D" },
                    DateReleased = DateTime.Now,
                };
                var videoGameUpdateValues = new VideoGame
                {
                    Name = "Vinkings on the Ocean 4",
                    Developers = new List<string> { "1", "2", "3" },
                    Publishers = new List<string> { "1", "2" },
                    Directors = new List<string> { "1" },
                    Producers = new List<string> { "1", "2" },
                    Designers = new List<string> { "1", "2", "3" },
                    Programmers = new List<string> { "1", "2" },
                    Artists = new List<string> { "1", "2", "3" },
                    Composers = new List<string> { "1", "2" },
                    Platforms = new List<string> { "1", "2" },
                    DateReleased = DateTime.Now,
                };
                var updatedVideoGame = new VideoGame();
                var expectedQuery = new Mock<IQuery<VideoGame>>();

                mockVideoGameStorage.Setup(m => m.Get(id)).Returns(videoGameToUpdate);
                mockVideoGameStorage.Setup(m => m.Update(videoGameToUpdate)).Returns(updatedVideoGame);
                mockQueryFactory.Setup(m => m.Build(default, default, updatedVideoGame)).Returns(expectedQuery.Object);

                // Act
                var actual = subject.Update(id, videoGameUpdateValues);

                // Assert
                actual.ShouldBe(expectedQuery.Object);

                mockVideoGameStorage.Verify(
                    m => m.Update(It.Is<VideoGame>(q =>
                    q.Name == videoGameUpdateValues.Name &&
                    q.Developers == videoGameUpdateValues.Developers &&
                    q.Publishers == videoGameUpdateValues.Publishers &&
                    q.Directors == videoGameUpdateValues.Directors &&
                    q.Producers == videoGameUpdateValues.Producers &&
                    q.Designers == videoGameUpdateValues.Designers &&
                    q.Programmers == videoGameUpdateValues.Programmers &&
                    q.Artists == videoGameUpdateValues.Artists &&
                    q.Composers == videoGameUpdateValues.Composers &&
                    q.Platforms == videoGameUpdateValues.Platforms &&
                    q.DateReleased == videoGameUpdateValues.DateReleased)), Times.Once);

                mockPrepVideoGame.Verify(m => m.Invoke(updatedVideoGame));
            }

            [Test]
            public void UpdateVideoGameInStorageWithSomeSetFields()
            {
                // Arrange
                var id = 5u;
                var videoGameToUpdate = new VideoGame
                {
                    Id = id,
                    Name = "Vikings at Sea IV",
                    Developers = new List<string> { "A", "B" },
                    Publishers = new List<string> { "A", "B", "C" },
                    Directors = new List<string> { "A" },
                    Producers = new List<string> { "A" },
                    Designers = new List<string> { "A", "B" },
                    Programmers = new List<string> { "A", "B", "C", "D" },
                    Artists = new List<string> { "A", "B" },
                    Composers = new List<string> { "A", "B", "C" },
                    Platforms = new List<string> { "A", "B", "C", "D" },
                    DateReleased = DateTime.Now,
                };
                var videoGameUpdateValues = new VideoGame
                {
                    Name = string.Empty,
                    Publishers = new List<string> { "1", "2" },
                    Directors = new List<string> { "1" },
                    Designers = new List<string>(),
                    Artists = new List<string> { "1", "2", "3" },
                    Composers = new List<string> { "1", "2" },
                };
                var updatedVideoGame = new VideoGame();
                var expectedQuery = new Mock<IQuery<VideoGame>>();

                mockVideoGameStorage.Setup(m => m.Get(id)).Returns(videoGameToUpdate);
                mockVideoGameStorage.Setup(m => m.Update(videoGameToUpdate)).Returns(updatedVideoGame);
                mockQueryFactory.Setup(m => m.Build(default, default, updatedVideoGame)).Returns(expectedQuery.Object);

                // Act
                var actual = subject.Update(id, videoGameUpdateValues);

                // Assert
                actual.ShouldBe(expectedQuery.Object);

                mockVideoGameStorage.Verify(
                    m => m.Update(It.Is<VideoGame>(q =>
                    q.Id != videoGameUpdateValues.Id &&
                    q.Name != videoGameUpdateValues.Name &&
                    q.Developers != videoGameUpdateValues.Developers &&
                    q.Publishers == videoGameUpdateValues.Publishers &&
                    q.Directors == videoGameUpdateValues.Directors &&
                    q.Producers != videoGameUpdateValues.Producers &&
                    q.Designers == videoGameUpdateValues.Designers &&
                    q.Programmers != videoGameUpdateValues.Programmers &&
                    q.Artists == videoGameUpdateValues.Artists &&
                    q.Composers == videoGameUpdateValues.Composers &&
                    q.Platforms != videoGameUpdateValues.Platforms &&
                    q.DateReleased != videoGameUpdateValues.DateReleased)), Times.Once);

                mockPrepVideoGame.Verify(m => m.Invoke(updatedVideoGame));
            }

            [Test]
            public void ReturnUnsuccessfulQueryWhenVideoGameNotFound()
            {
                // Arrange
                var id = 99u;
                var expectedQuery = new Mock<IQuery<VideoGame>>();

                mockQueryFactory.Setup(m => m.Build<VideoGame>(404, ResponseMessages.VideoGame.NotFound(id), default)).Returns(expectedQuery.Object);

                // Act
                var actual = subject.Update(id, new VideoGame());

                // Assert
                actual.ShouldBe(expectedQuery.Object);
            }
        }

        private class DeleteShould : VideoGameServiceTests
        {
            [Test]
            public void DeleteVideoGameInStorageAndReturnQuery()
            {
                // Arrange
                var id = 5u;
                var deletedVideoGame = new VideoGame();
                var expectedQuery = new Mock<IQuery<VideoGame>>();

                mockVideoGameStorage.Setup(m => m.Delete(id)).Returns(deletedVideoGame);
                mockQueryFactory.Setup(m => m.Build(default, ResponseMessages.VideoGame.Deleted(id), deletedVideoGame)).Returns(expectedQuery.Object);

                // Act
                var actual = subject.Delete(id);

                // Assert
                actual.ShouldBe(expectedQuery.Object);
            }

            [Test]
            public void ReturnUnsuccessfulQueryWhenVideoGameNotFound()
            {
                // Arrange
                var id = 99u;
                var expectedQuery = new Mock<IQuery<VideoGame>>();

                mockQueryFactory.Setup(m => m.Build<VideoGame>(404, ResponseMessages.VideoGame.NotFound(id), default)).Returns(expectedQuery.Object);

                // Act
                var actual = subject.Delete(id);

                // Assert
                actual.ShouldBe(expectedQuery.Object);
            }
        }
    }
}
