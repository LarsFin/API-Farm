using System.Collections.Generic;
using ApiFarm.Models;
using ApiFarm.Repositories;
using ApiFarm.Utils;
using Moq;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Repositories
{
    public class InMemoryTests
    {
        private InMemory<Model> subject;

        private Mock<ICloner<Model>> mockCloner;
        private List<Model> models;

        [SetUp]
        public void Setup()
        {
            mockCloner = new Mock<ICloner<Model>>();
            models = new List<Model>();
            subject = new InMemory<Model>(mockCloner.Object, models);
        }

        public class Model : IModel
        {
            public uint Id { get; set; }
        }

        private class GetAllShould : InMemoryTests
        {
            [Test]
            public void ReturnEmptyModelsList()
            {
                // Arrange
                subject = new InMemory<Model>(mockCloner.Object);

                // Act
                var retrievedModels = subject.GetAll();

                // Assert
                retrievedModels.ShouldBeEmpty();
            }

            [Test]
            public void ReturnListOfClonedModels()
            {
                // Arrange
                var originalModel = new Model();
                models.Add(originalModel);
                var cloneModel = new Model();

                mockCloner.Setup(m => m.Clone(originalModel)).Returns(cloneModel);

                // Act
                var retrievedModels = subject.GetAll();

                // Assert
                retrievedModels.ShouldNotBeSameAs(models);
                retrievedModels.ShouldContain(cloneModel);
            }
        }
    }
}
