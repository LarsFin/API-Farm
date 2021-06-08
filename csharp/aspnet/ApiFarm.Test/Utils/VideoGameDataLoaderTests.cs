using System;
using System.Collections.Generic;
using ApiFarm.Models.Impl;
using ApiFarm.Utils.Impl;
using Moq;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Utils
{
    public class VideoGameDataLoaderTests
    {
        private VideoGameDataLoader subject;

        private Mock<Func<string, string>> stubbedReadFile;
        private Mock<Func<string, IEnumerable<VideoGame>>> stubbedParseJson;

        [SetUp]
        public void SetUp()
        {
            stubbedReadFile = new Mock<Func<string, string>>();
            stubbedParseJson = new Mock<Func<string, IEnumerable<VideoGame>>>();

            subject = new VideoGameDataLoader(stubbedReadFile.Object, stubbedParseJson.Object);
        }

        private class LoadShould : VideoGameDataLoaderTests
        {
            [Test]
            public void ReturnVideoGames()
            {
                // Arrange
                var dataPath = "PATH TO SAMPLE DATA FILE";
                var jsonSampleData = "JSON SAMPLE DATA";
                var expected = new Mock<IEnumerable<VideoGame>>();

                stubbedReadFile.Setup(m => m.Invoke(dataPath)).Returns(jsonSampleData);
                stubbedParseJson.Setup(m => m.Invoke(jsonSampleData)).Returns(expected.Object);

                // Act
                var actual = subject.Load(dataPath);

                // Assert
                actual.ShouldBeSameAs(expected.Object);
            }
        }
    }
}
