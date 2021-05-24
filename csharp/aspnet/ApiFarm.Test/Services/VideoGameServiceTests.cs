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
                var expectedCode = 200u;

                var storedVideoGames = new Mock<IEnumerable<VideoGame>>();
                var expected = new Mock<IQuery<IEnumerable<VideoGame>>>();

                mockVideoGameStorage.Setup(m => m.GetAll()).Returns(storedVideoGames.Object);
                mockQueryFactory.Setup(m => m.Build<IEnumerable<VideoGame>>(expectedCode, default, storedVideoGames.Object))
                    .Returns(expected.Object);

                // Act
                var actual = subject.GetAll();

                // Assert
                actual.ShouldBe(expected.Object);
            }
        }
    }
}
