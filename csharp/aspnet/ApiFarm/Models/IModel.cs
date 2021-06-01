namespace ApiFarm.Models
{
    /// <summary>
    /// Describes required behaviours of a model which can be stored with ApiFarm facilitators.
    /// </summary>
    public interface IModel
    {
        /// <summary>
        /// Gets or sets identifier of model.
        /// </summary>
        uint Id { get; set; }
    }
}
