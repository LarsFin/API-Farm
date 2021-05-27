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

        [SetUp]
        protected void SetUp()
        {
            mockVideoGameStorage = new Mock<IRepository<VideoGame>>();
            mockQueryFactory = new Mock<IQueryFactory>();

            subject = new VideoGameService(
                mockVideoGameStorage.Object,
                mockQueryFactory.Object);
        }

        private class GetAllShould : VideoGameServiceTests
        {
            [Test]
            public void RetrieveVideoGamesFromStorageAndReturnQuery()
            {
                // Arrange
                var storedVideoGames = new Mock<IEnumerable<VideoGame>>();
                var expected = new Mock<IQuery<IEnumerable<VideoGame>>>();

                mockVideoGameStorage.Setup(m => m.GetAll()).Returns(storedVideoGames.Object);
                mockQueryFactory.Setup(m => m.Build(default, default, storedVideoGames.Object))
                    .Returns(expected.Object);

                // Act
                var actual = subject.GetAll();

                // Assert
                actual.ShouldBe(expected.Object);
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

                mockQueryFactory.Setup(m => m.Build(400, "A name is required for a video game.", default(VideoGame))).Returns(expectedQuery.Object);

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

                mockQueryFactory.Setup(m => m.Build(400, "A date_released is required for a video game.", default(VideoGame))).Returns(expectedQuery.Object);

                // Act
                var actual = subject.Add(videoGame);

                // Assert
                actual.ShouldBe(expectedQuery.Object);
            }
        }
    }
}
